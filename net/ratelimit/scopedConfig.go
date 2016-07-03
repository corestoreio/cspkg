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

package ratelimit

import (
	"net/http"

	"github.com/corestoreio/csfw/store/scope"
	"github.com/corestoreio/csfw/util/errors"
	"gopkg.in/throttled/throttled.v2"
)

// scopedConfig private internal scoped based configuration
type scopedConfig struct {
	scopedConfigGeneric

	// start of package specific config values

	// disabled set to true to disable rate limiting
	disabled bool
	// deniedHandler can be customized instead of showing a HTTP status 429
	// error page once the HTTPRateLimit has been reached.
	// It will be called if the request gets over the limit.
	deniedHandler http.Handler

	// RateLimiter default not set. It gets set either through the developer
	// calling WithRateLimiter() or via OptionFactoryFunc.
	throttled.RateLimiter

	// VaryByer is called for each request to generate a key for the limiter. If
	// it is nil, the middleware panics. The default VaryByer returns an empty
	// string so that all requests uses the same key.
	VaryByer
}

var defaultDeniedHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
})

func defaultScopedConfig(h scope.Hash) *scopedConfig {
	return &scopedConfig{
		scopedConfigGeneric: scopedConfigGeneric{
			scopeHash: h,
		},
		deniedHandler: defaultDeniedHandler,
		VaryByer:      emptyVaryBy{},
	}
}

// isValid a configuration for a scope is only then valid when several fields
// are not empty and scopedConfig itself has a valid pointer.
func (sc *scopedConfig) isValid() error {
	if sc == nil {
		return errors.NewNotValidf(errScopedConfigIsNil)
	}

	if sc.lastErr != nil {
		return errors.Wrap(sc.lastErr, "[ratelimit] scopedConfig.isValid has an lastErr")
	}
	if sc.scopeHash == 0 || sc.RateLimiter == nil || sc.deniedHandler == nil || sc.VaryByer == nil {
		return errors.NewNotValidf(errScopedConfigNotValid, sc.scopeHash, sc.deniedHandler == nil, sc.RateLimiter == nil, sc.VaryByer == nil)
	}
	return nil
}

func (sc *scopedConfig) requestRateLimit(r *http.Request) (bool, throttled.RateLimitResult, error) {
	return sc.RateLimiter.RateLimit(sc.VaryByer.Key(r), 1)
}