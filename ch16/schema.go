package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

type Person struct {
	Name  string
	Phone string
}

func main() {
	http.HandleFunc("/schema", func(res http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(0)
		var person Person
		decoder.Decode(&person, req.PostForm)
		message := fmt.Sprintf("Hello %v from area %v", person.Name, person.Phone)
		fmt.Printf("%v", req.PostForm)
		res.Write([]byte(message))
	})

	http.ListenAndServe(":8080", nil)
}
