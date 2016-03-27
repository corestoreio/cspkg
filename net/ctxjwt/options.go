// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ctxjwt

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"crypto/rand"

	"github.com/corestoreio/csfw/config"
	"github.com/corestoreio/csfw/net/ctxhttp"
	"github.com/corestoreio/csfw/store/scope"
	"github.com/dgrijalva/jwt-go"
	"github.com/juju/errors"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

type scopedConfig struct {
	rsapk        *rsa.PrivateKey
	ecdsapk      *ecdsa.PrivateKey
	hmacPassword []byte // password for hmac
	hasKey       bool   // must be set to true if one of the three above keys has been set

	// expire defines the duration when the token is about to expire
	expire time.Duration
	// signingMethod how to sign the JWT. For default value see the OptionFuncs
	signingMethod jwt.SigningMethod
	// enableJTI activates the (JWT ID) Claim, a unique identifier. UUID.
	enableJTI    bool
	errorHandler ctxhttp.HandlerFunc
	// keyFunc will receive the parsed token and should return the key for validating.
	keyFunc jwt.Keyfunc
}

// ErrPrivateKeyNotFound will be returned when the PK cannot be read from the Reader
var ErrPrivateKeyNotFound = errors.New("Private Key from io.Reader no found")

// ErrPrivateKeyNoPassword will be returned when the PK is encrypted but you
// forgot to provide a password.
var ErrPrivateKeyNoPassword = errors.New("Private Key is encrypted but password was not set")

// PrivateKeyBits used when auto generating a private key
const PrivateKeyBits = 4096

// @todo add more WithKeyFrom...()

// Option can be used as an argument in NewService to configure a token service.
type Option func(*Service)

// WithDefaultConfig applies the default JWT configuration settings based for
// a specific scope.
//
// Default values are:
//		- constant DefaultExpire
//		- HMAC Password: uuid.NewRandom()
//		- Signing Method HMAC SHA 256
//		- HTTP error handler returns http.StatusUnauthorized
func WithDefaultConfig(scp scope.Scope, id int64) Option {
	h := scope.NewHash(scp, id)
	return func(s *Service) {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.scopeCache[h] = scopedConfig{
			expire:        DefaultExpire,
			hmacPassword:  []byte(uuid.NewRandom()), // @todo can be better ...
			signingMethod: jwt.SigningMethodHS256,
			enableJTI:     false,
			errorHandler: ctxhttp.HandlerFunc(func(_ context.Context, w http.ResponseWriter, _ *http.Request) error {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return nil
			}),
		}
	}
}

// WithBlacklist sets a new global black list service.
func WithBlacklist(blacklist Blacklister) Option {
	return func(s *Service) {
		s.Blacklist = blacklist
	}
}

// WithBackend applies the backend configuration to the service.
// Once this has been set all other option functions are not really
// needed.
func WithBackend(pb *PkgBackend) Option {
	return func(s *Service) {
		s.backend = pb
	}
}

// WithPassword sets the HMAC 256 bit signing method with a password.
// Useful to use Magento encryption key as the key.
func WithPassword(scp scope.Scope, id int64, key []byte) Option {
	h := scope.NewHash(scp, id)
	return func(s *Service) {
		s.mu.Lock()
		defer s.mu.Unlock()

		scNew := scopedConfig{
			hasKey:        true,
			signingMethod: jwt.SigningMethodHS256,
			hmacPassword:  key,
		}
		if sc, ok := s.scopeCache[h]; ok {
			sc.hasKey = scNew.hasKey
			sc.signingMethod = scNew.signingMethod
			sc.hmacPassword = scNew.hmacPassword
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// WithSigningMethod this option function lets you overwrite the default 256 bit
// signing method for a specific scope. Used incorrectly token decryption can fail.
func WithSigningMethod(scp scope.Scope, id int64, sm jwt.SigningMethod) Option {
	h := scope.NewHash(scp, id)
	return func(s *Service) {
		s.mu.Lock()
		defer s.mu.Unlock()

		scNew := scopedConfig{
			signingMethod: sm,
		}
		if sc, ok := s.scopeCache[h]; ok {
			sc.signingMethod = sm
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// WithErrorHandler sets the error handler for a scope and its ID.
func WithErrorHandler(scp scope.Scope, id int64, handler ctxhttp.HandlerFunc) Option {
	h := scope.NewHash(scp, id)
	return func(s *Service) {
		s.mu.Lock()
		defer s.mu.Unlock()
		scNew := scopedConfig{
			errorHandler: handler,
		}
		if sc, ok := s.scopeCache[h]; ok {
			sc.errorHandler = scNew.errorHandler
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// WithExpiration sets expiration duration depending on the scope
func WithExpiration(scp scope.Scope, id int64, d time.Duration) Option {
	h := scope.NewHash(scp, id)
	return func(s *Service) {
		s.mu.Lock()
		defer s.mu.Unlock()

		scNew := scopedConfig{
			expire: d,
		}

		if sc, ok := s.scopeCache[h]; ok {
			sc.expire = scNew.expire
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// WithTokenID enables JTI (JSON Web Token ID) for a specific scope
func WithTokenID(scp scope.Scope, id int64, enable bool) Option {
	h := scope.NewHash(scp, id)
	return func(s *Service) {
		s.mu.Lock()
		defer s.mu.Unlock()

		scNew := scopedConfig{
			enableJTI: enable,
		}

		if sc, ok := s.scopeCache[h]; ok {
			sc.enableJTI = scNew.enableJTI
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// WithECDSAFromFile loads the ECDSA key from a file @todo
func WithECDSAFromFile(scp scope.Scope, id int64, fileName string, password ...[]byte) Option {
	fpk, err := ioutil.ReadFile(fileName)
	if err != nil {
		return func(s *Service) {
			s.MultiErr = s.AppendErrors(err)
		}
	}
	return WithECDSA(scp, id, fpk, password...)
}

// WithECDSA @todo
// Default Signing bits 256.
func WithECDSA(scp scope.Scope, id int64, privateKey []byte, password ...[]byte) Option {
	err := errors.New("WithECDSA: TODO: implement")
	return withECDSA(scope.NewHash(scp, id), nil, err)
}

func withECDSA(h scope.Hash, ecdsapk *ecdsa.PrivateKey, err error) Option {
	return func(s *Service) {
		if err != nil {
			s.MultiErr = s.AppendErrors(err)
			return
		}

		s.mu.Lock()
		defer s.mu.Unlock()

		scNew := scopedConfig{
			hasKey:        true,
			signingMethod: jwt.SigningMethodES256,
			ecdsapk:       ecdsapk,
		}

		if sc, ok := s.scopeCache[h]; ok {
			sc.hasKey = scNew.hasKey
			sc.signingMethod = scNew.signingMethod
			sc.rsapk = scNew.rsapk
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// WithRSAFromFile reads an RSA private key from a file and applies it as an option
// to the AuthManager. Password as second argument is only required when the
// private key is encrypted. Public key will be derived from the private key.
func WithRSAFromFile(scp scope.Scope, id int64, fileName string, password ...[]byte) Option {
	fpk, err := ioutil.ReadFile(fileName)
	if err != nil {
		return func(s *Service) {
			s.MultiErr = s.AppendErrors(err)
		}
	}
	return WithRSA(scp, id, fpk, password...)
}

// WithRSA reads PEM byte data and decodes it and parses the private key.
// Applies the private and the public key to the AuthManager. Password as second
// argument is only required when the private key is encrypted.
// Checks for io.Close and closes the resource. Public key will be derived from
// the private key. Default Signing bits 256.
func WithRSA(scp scope.Scope, id int64, privateKey []byte, password ...[]byte) Option {

	var prKeyPEM *pem.Block
	if prKeyPEM, _ = pem.Decode(privateKey); prKeyPEM == nil {
		return func(s *Service) {
			s.MultiErr = s.AppendErrors(ErrPrivateKeyNotFound)
		}
	}

	var rsaPrivateKey *rsa.PrivateKey
	var err error
	if x509.IsEncryptedPEMBlock(prKeyPEM) {
		if len(password) != 1 || len(password[0]) == 0 {
			return func(s *Service) {
				if PkgLog.IsDebug() {
					PkgLog.Debug("ctxjwt.WithRSA.IsEncryptedPEMBlock", "err", ErrPrivateKeyNoPassword)
				}
				s.MultiErr = s.AppendErrors(ErrPrivateKeyNoPassword)
			}
		}
		var dd []byte
		var errPEM error
		if dd, errPEM = x509.DecryptPEMBlock(prKeyPEM, password[0]); errPEM != nil {
			return func(s *Service) {
				if PkgLog.IsDebug() {
					PkgLog.Debug("ctxjwt.WithRSA.DecryptPEMBlock", "err", errPEM)
				}
				s.MultiErr = s.AppendErrors(errPEM)
			}
		}
		rsaPrivateKey, err = x509.ParsePKCS1PrivateKey(dd)
	} else {
		rsaPrivateKey, err = x509.ParsePKCS1PrivateKey(prKeyPEM.Bytes)
	}

	return withRSA(scope.NewHash(scp, id), rsaPrivateKey, err)
}

// WithRSAGenerator creates an in-memory RSA key without persisting it.
// This function may run around ~3secs.
func WithRSAGenerator(scp scope.Scope, id int64) Option {
	pk, err := rsa.GenerateKey(rand.Reader, PrivateKeyBits)
	return withRSA(scope.NewHash(scp, id), pk, err)
}

// withRSA internal option functions which adds a RSA private key to the Service
func withRSA(h scope.Hash, pk *rsa.PrivateKey, err error) Option {
	return func(s *Service) {
		if err != nil {
			s.MultiErr = s.AppendErrors(err)
			return
		}

		s.mu.Lock()
		defer s.mu.Unlock()

		scNew := scopedConfig{
			hasKey:        true,
			signingMethod: jwt.SigningMethodRS256,
			rsapk:         pk,
		}

		if sc, ok := s.scopeCache[h]; ok {
			sc.hasKey = scNew.hasKey
			sc.signingMethod = scNew.signingMethod
			sc.rsapk = scNew.rsapk
			scNew = sc
		}
		s.scopeCache[h] = scNew
	}
}

// optionsByBackend creates an option array containing the Options based
// on the configuration
func optionsByBackend(be *PkgBackend, sg config.ScopedGetter) (opts [5]Option) {
	var emptySC scopedConfig
	scp, id := sg.Scope()

	if val, err := be.NetCtxjwtExpiration.Get(sg); err != nil {
		return emptySC, errors.Mask(err)
	} else {
		// golint is wrong here because we make use of the block scope in an if/else block
		// and also we don't want to spread the val variable all over the place.
		opts[0] = WithExpiration(scp, id, val)
	}

	if val, err := be.NetCtxjwtEnableJTI.Get(sg); err != nil {
		return emptySC, errors.Mask(err)
	} else {
		// golint is wrong here because we make use of the block scope in an if/else block
		// and also we don't want to spread the val variable all over the place.
		opts[1] = WithTokenID(scp, id, val)
	}

	sm, err := be.NetCtxjwtSigningMethod.Get(sg)
	if err != nil {
		return emptySC, errors.Mask(err)
	}

	switch sm.Alg() {
	case jwt.SigningMethodRS256.Alg(), jwt.SigningMethodRS384.Alg(), jwt.SigningMethodRS512.Alg():

		rsaKey, err := be.NetCtxjwtRSAKey.Get(sg)
		if err != nil {
			return emptySC, errors.Mask(err)
		}
		rsaPassword, err := be.NetCtxjwtRSAKeyPassword.Get(sg)
		if err != nil {
			return emptySC, errors.Mask(err)
		}

		opts[2] = WithRSA(scp, id, rsaKey, rsaPassword)

	case jwt.SigningMethodES256.Alg(), jwt.SigningMethodES384.Alg(), jwt.SigningMethodES512.Alg():

		ecdsaKey, err := be.NetCtxjwtECDSAKey.Get(sg)
		if err != nil {
			return emptySC, errors.Mask(err)
		}
		ecdsaPassword, err := be.NetCtxjwtECDSAKeyPassword.Get(sg)
		if err != nil {
			return emptySC, errors.Mask(err)
		}

		opts[3] = WithECDSA(scp, id, ecdsaKey, ecdsaPassword)

	case jwt.SigningMethodHS256.Alg(), jwt.SigningMethodHS384.Alg(), jwt.SigningMethodHS512.Alg():

		password, err := be.NetCtxjwtHmacPassword.Get(sg)
		if err != nil {
			return emptySC, errors.Mask(err)
		}

		opts[4] = WithPassword(scp, id, password)

	default:
		return emptySC, ErrUnexpectedSigningMethod
	}
	return opts
}
