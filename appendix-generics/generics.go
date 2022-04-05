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

type Animal interface {
	Sound()
}

type Cat struct{}

func (c Cat) Sound() { fmt.Println("Meow") }

type Dog struct{}

func (d Dog) Sound() { fmt.Println("Woof") }

type Hybrid interface {
	Animal
}

func test[H Hybrid](which H) H {
	which.Sound()
	return which
	/*
		switch any(which).(type) {
		case Dog:
			which.Sound()
			return which
		case Cat:
			which.Sound()
			return which
		default:
			which.Sound()
			return which
		}
	*/
}

type myStruct struct {
	s *strings.Reader
}

func (m myStruct) Len() int {
	return m.s.Len()
}

func (m myStruct) Read(b []byte) (int, error) {
	return m.s.Read(b)
}

func main() {
	c := test(Cat{})
	d := test(Dog{})

	c = c
	d = d

	r1 := NewRequest("GET", "/", myStruct{strings.NewReader("")})
	r2 := NewRequest("GET", "/", myStruct{})

	r1 = r1
	r2 = r2
}

type MyType interface {
	bytes.Buffer | bytes.Reader | strings.Reader | myStruct

	Len() int
	io.Reader
	comparable
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
		req.ContentLength = int64(body.Len())
		switch v := any(body).(type) {
		case io.ReadCloser:
			req.Body = v
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
