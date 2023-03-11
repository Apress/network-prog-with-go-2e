package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func buildHandler(message string) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", buildHandler("HomeHandler"))
	r.HandleFunc("/products", buildHandler("ProductsHandler"))
	l := handlers.ContentTypeHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("json articles only"))
	}), "application/json")
	r.Handle("/articles", l)
	http.ListenAndServe(":8080", r)
}
