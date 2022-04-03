package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FuzzMe(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}

	// the fuzz runner f leverages the test runner t,
	// this is so the fuzzer can manage the tests, it generates (or uses seed) inputs
	// calling the passed in test
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := hex.EncodeToString(in)
		out, err := hex.DecodeString(enc)
		if err != nil {
			t.Fatalf("%v: decode: %v", in, err)
		}
		if !bytes.Equal(in, out) {
			t.Fatalf("%v: not equal after round trip: %v", in, out)
		}
	})
}

func FuzzBad(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		if i != i {
			f.Fatalf("want: %v, got: %v", i, i)
		}
	})
}

func FuzzHandler(f *testing.F) {
	f.Fuzz(func(t *testing.T, data string) {
		v := base64.StdEncoding.EncodeToString([]byte(data))
		req := httptest.NewRequest("GET", "/?q="+v, nil)
		res := httptest.NewRecorder()

		f := func(w http.ResponseWriter, req *http.Request) {
			keys, ok := req.URL.Query()["q"]

			if !ok || len(keys) != 1 {
				t.Log(keys)
				t.Fatal("q param missing or more than one instance")
			}

			val := keys[0]

			if len(val) > 16384 {
				w.WriteHeader(http.StatusNotAcceptable)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}

		f(res, req)

		if res == nil || res.Result().StatusCode != http.StatusOK {
			t.Fatal(res)
		}
	})
}
