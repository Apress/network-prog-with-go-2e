package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

var hashKey = []byte("very-secret")
var s = securecookie.New(hashKey, nil)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("cookie-name", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)

		fmt.Println(cookie)
	}
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)
		fmt.Println(cookie.Value)
		if err = s.Decode("cookie-name", cookie.Value, &value); err == nil {
			fmt.Println("im ok!")
			fmt.Fprintf(w, "The value of foo is %q", value["foo"])
		}

		fmt.Println(cookie, err)
	}

}

func main() {
	http.HandleFunc("/set", SetCookieHandler)
	http.HandleFunc("/read", ReadCookieHandler)
	http.ListenAndServe(":8080", nil)
}
