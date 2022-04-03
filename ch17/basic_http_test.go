package ch17

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPRoundTripNoConnection(t *testing.T) {
	path := "/"
	req := httptest.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()

	f := func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(5 * time.Second) // holding connection
	}

	f(res, req)

	if res == nil || res.Result().StatusCode != http.StatusOK {
		t.Error(res)
	}
}

func TestHTTPRoundTrip(t *testing.T) {
	path := "/"
	c := make(chan struct{})
	//server
	go func() {
		http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			time.Sleep(5 * time.Second) // holding connection
		})
		http.ListenAndServe(":8080", nil)
	}()

	//client
	go func() {
		resp, err := http.Get("http://localhost:8080" + path)
		if err != nil {
			t.Error(err)
		} else {
			if resp == nil || resp.StatusCode != http.StatusOK {
				t.Error(resp)
			}
		}

		defer func() {
			c <- struct{}{}
		}()
	}()
	<-c
}

func TestHTTPTestRoundTripTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(5 * time.Second) // holding connection
	}))
	defer ts.Close()

	var c = &http.Client{
		Timeout: time.Second * 2,
	}
	req, _ := http.NewRequest("GET", ts.URL, nil)
	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	err = res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

}

func TestHTTPRoundTripTimeout(t *testing.T) {
	path := "/unique"
	c := make(chan struct{})
	//server
	go func() {
		http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			time.Sleep(5 * time.Second) // holding connection
		})
		http.ListenAndServe(":8080", nil)
	}()

	//client
	go func() {
		var client = &http.Client{
			Timeout: time.Second * 2,
		}
		resp, err := client.Get("http://localhost:8080" + path)
		if err != nil {
			t.Error(err)
		} else {
			if resp == nil || resp.StatusCode != http.StatusOK {
				t.Error(resp)
			}
		}

		defer func() {
			c <- struct{}{}
		}()
	}()
	<-c
}

func TestPipe(t *testing.T) {
	c := make(chan struct{})
	server, client := net.Pipe()
	go func() {
		time.Sleep(2 * time.Second)
		req := make([]byte, 15)
		server.SetDeadline(time.Now().Add(1 * time.Second))
		_, err := server.Read(req)
		t.Log(string(req))
		if err != nil {
			t.Error(err)
		}
		defer func() {
			server.Close()
			c <- struct{}{}
		}()
	}()
	client.SetDeadline(time.Now().Add(1 * time.Second))
	_, err := client.Write([]byte("my http request"))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		client.Close()
	}()
	<-c
}
