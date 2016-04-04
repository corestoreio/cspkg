package csjwt_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/corestoreio/csfw/util/csjwt"
)

// Sample data from http://tools.ietf.org/html/draft-jones-json-web-signature-04#appendix-A.1
var hmacTestKey []byte

const HMACFastSuffix csjwt.Algorithm = 10

func init() {
	var err error
	hmacTestKey, err = ioutil.ReadFile("test/hmacTestKey")
	if err != nil {
		panic(err)
	}

	hf, err := csjwt.NewHMACFast256(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		panic(err)
	}
	hf.Name *= HMACFastSuffix
	csjwt.RegisterSigningMethod(hf)

	hf, err = csjwt.NewHMACFast384(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		panic(err)
	}
	hf.Name *= HMACFastSuffix
	csjwt.RegisterSigningMethod(hf)

	hf, err = csjwt.NewHMACFast512(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		panic(err)
	}
	hf.Name *= HMACFastSuffix
	csjwt.RegisterSigningMethod(hf)
}

var hmacTestData = []struct {
	name        string
	tokenString []byte
	alg         string
	claims      map[string]interface{}
	valid       bool
}{
	{
		"web sample",
		[]byte("eyJ0eXAiOiJKV1QiLA0KICJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ.dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"),
		"HS256",
		map[string]interface{}{"iss": "joe", "exp": 1300819380, "http://example.com/is_root": true},
		true,
	},
	{
		"HS384",
		[]byte("eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJleHAiOjEuMzAwODE5MzhlKzA5LCJodHRwOi8vZXhhbXBsZS5jb20vaXNfcm9vdCI6dHJ1ZSwiaXNzIjoiam9lIn0.KWZEuOD5lbBxZ34g7F-SlVLAQ_r5KApWNWlZIIMyQVz5Zs58a7XdNzj5_0EcNoOy"),
		"HS384",
		map[string]interface{}{"iss": "joe", "exp": 1300819380, "http://example.com/is_root": true},
		true,
	},
	{
		"HS512",
		[]byte("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEuMzAwODE5MzhlKzA5LCJodHRwOi8vZXhhbXBsZS5jb20vaXNfcm9vdCI6dHJ1ZSwiaXNzIjoiam9lIn0.CN7YijRX6Aw1n2jyI2Id1w90ja-DEMYiWixhYCyHnrZ1VfJRaFQz1bEbjjA5Fn4CLYaUG432dEYmSbS4Saokmw"),
		"HS512",
		map[string]interface{}{"iss": "joe", "exp": 1300819380, "http://example.com/is_root": true},
		true,
	},
	{
		"web sample: invalid",
		[]byte("eyJ0eXAiOiJKV1QiLA0KICJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ.dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXo"),
		"HS256",
		map[string]interface{}{"iss": "joe", "exp": 1300819380, "http://example.com/is_root": true},
		false,
	},
}

func TestHMACVerify(t *testing.T) {
	for _, data := range hmacTestData {
		signing, signature, err := csjwt.SplitForVerify(data.tokenString)
		if err != nil {
			t.Fatal(err, "\n", string(data.tokenString))
		}

		method := csjwt.GetSigningMethod(data.alg)
		err = method.Verify(signing, signature, csjwt.WithPassword(hmacTestKey))
		if data.valid && err != nil {
			t.Errorf("[%v] Method %s Error while verifying key: %v", data.name, data.alg, err)
		}
		if !data.valid && err == nil {
			t.Errorf("[%v] Invalid key passed validation", data.name)
		}
	}
}

func TestHMACSign(t *testing.T) {
	for _, data := range hmacTestData {
		if data.valid {
			signing, signature, err := csjwt.SplitForVerify(data.tokenString)
			if err != nil {
				t.Fatal(err, "\n", string(data.tokenString))
			}

			method := csjwt.GetSigningMethod(data.alg)
			sig, err := method.Sign(signing, csjwt.WithPassword(hmacTestKey))
			if err != nil {
				t.Errorf("[%v] Error signing token: %v", data.name, err)
			}
			if !bytes.Equal(sig, signature) {
				t.Errorf("[%v] Incorrect signature.\nwas:\n%v\nexpecting:\n%v", data.name, string(sig), string(signature))
			}
		}
	}
}

func BenchmarkHS256Signing(b *testing.B) {
	benchmarkSigning(b, csjwt.SigningMethodHS256, csjwt.WithPassword(hmacTestKey))
}

func BenchmarkHS384Signing(b *testing.B) {
	benchmarkSigning(b, csjwt.SigningMethodHS384, csjwt.WithPassword(hmacTestKey))
}

func BenchmarkHS512Signing(b *testing.B) {
	benchmarkSigning(b, csjwt.SigningMethodHS512, csjwt.WithPassword(hmacTestKey))
}

func BenchmarkHS256SigningFast(b *testing.B) {
	hf, err := csjwt.NewHMACFast256(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		b.Fatal(err)
	}
	benchmarkSigning(b, hf, csjwt.Key{})
}

func BenchmarkHS384SigningFast(b *testing.B) {
	hf, err := csjwt.NewHMACFast384(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		b.Fatal(err)
	}
	benchmarkSigning(b, hf, csjwt.Key{})
}

func BenchmarkHS512SigningFast(b *testing.B) {
	hf, err := csjwt.NewHMACFast512(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		b.Fatal(err)
	}
	benchmarkSigning(b, hf, csjwt.Key{})
}

func BenchmarkHS256Verify(b *testing.B) {
	signing, signature, err := csjwt.SplitForVerify(hmacTestData[0].tokenString) // HS256 token
	if err != nil {
		b.Fatal(err)
	}
	benchmarkMethodVerify(b, csjwt.SigningMethodHS256, signing, signature, csjwt.WithPassword(hmacTestKey))
}
func BenchmarkHS384Verify(b *testing.B) {
	signing, signature, err := csjwt.SplitForVerify(hmacTestData[1].tokenString) // HS384 token
	if err != nil {
		b.Fatal(err)
	}
	benchmarkMethodVerify(b, csjwt.SigningMethodHS384, signing, signature, csjwt.WithPassword(hmacTestKey))
}
func BenchmarkHS512Verify(b *testing.B) {
	signing, signature, err := csjwt.SplitForVerify(hmacTestData[2].tokenString) // HS512 token
	if err != nil {
		b.Fatal(err)
	}
	benchmarkMethodVerify(b, csjwt.SigningMethodHS512, signing, signature, csjwt.WithPassword(hmacTestKey))
}

func BenchmarkHS256VerifyFast(b *testing.B) {
	signing, signature, err := csjwt.SplitForVerify(hmacTestData[0].tokenString) // HS256 token
	if err != nil {
		b.Fatal(err)
	}
	hf, err := csjwt.NewHMACFast256(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		b.Fatal(err)
	}
	benchmarkMethodVerify(b, hf, signing, signature, csjwt.Key{})
}
func BenchmarkHS384VerifyFast(b *testing.B) {
	signing, signature, err := csjwt.SplitForVerify(hmacTestData[1].tokenString) // HS384 token
	if err != nil {
		b.Fatal(err)
	}
	hf, err := csjwt.NewHMACFast384(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		b.Fatal(err)
	}
	benchmarkMethodVerify(b, hf, signing, signature, csjwt.Key{})
}
func BenchmarkHS512VerifyFast(b *testing.B) {
	signing, signature, err := csjwt.SplitForVerify(hmacTestData[2].tokenString) // HS512 token
	if err != nil {
		b.Fatal(err)
	}
	hf, err := csjwt.NewHMACFast512(csjwt.WithPassword(hmacTestKey))
	if err != nil {
		b.Fatal(err)
	}
	benchmarkMethodVerify(b, hf, signing, signature, csjwt.Key{})
}
