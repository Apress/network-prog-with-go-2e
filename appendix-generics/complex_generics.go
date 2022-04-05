package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type myStruct struct {
	s *strings.Reader
}

func (m myStruct) Len() int {
	return m.s.Len()
}

func (m myStruct) Read(b []byte) (int, error) {
	return m.s.Read(b)
}

type MyType interface {
	*bytes.Buffer | *bytes.Reader | *strings.Reader | myStruct

	Len() int
	io.Reader
	comparable
}

type Lener interface {
	Len() int
}

// ./http/httptest/httptest.go
func NewRequest[M MyType](method, target string, body M) *http.Request {
	if method == "" {
		method = "GET"
	}
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(method + " " + target + " HTTP/1.0\r\n\r\n")))
	if err != nil {
		panic("invalid NewRequest arguments; " + err.Error())
	}

	// HTTP/1.0 was used above to avoid needing a Host field. Change it to 1.1 here.
	req.Proto = "HTTP/1.1"
	req.ProtoMinor = 1
	req.Close = false

	var zero M
	if body != zero {
		switch i := any(body).(type) {
		case Lener, io.ReadCloser:
			if b, ok := i.(Lener); ok {
				req.ContentLength = int64(b.Len())
			}
			if rc, ok := i.(io.ReadCloser); ok {
				req.Body = rc
			}
		default:
			req.Body = io.NopCloser(body)
		}
	} else {
		req.ContentLength = -1
	}

	// 192.0.2.0/24 is "TEST-NET" in RFC 5737 for use solely in
	// documentation and example source code and should not be
	// used publicly.
	req.RemoteAddr = "192.0.2.1:1234"

	if req.Host == "" {
		req.Host = "example.com"
	}

	if strings.HasPrefix(target, "https://") {
		req.TLS = &tls.ConnectionState{
			Version:           tls.VersionTLS12,
			HandshakeComplete: true,
			ServerName:        req.Host,
		}
	}

	return req
}

func main() {
	fmt.Println(NewRequest("GET", "/", myStruct{strings.NewReader("")}).ContentLength)
	fmt.Println(NewRequest("GET", "/", myStruct{}).ContentLength)
	fmt.Println(NewRequest("GET", "/", strings.NewReader("")).ContentLength)
	fmt.Println(NewRequest("GET", "/", &bytes.Buffer{}).ContentLength)
	fmt.Println(NewRequest("GET", "/", bytes.NewReader([]byte("read me"))).ContentLength)
}
