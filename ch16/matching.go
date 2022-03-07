package main

import (
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
	r.HandleFunc("/articles", buildHandler("ArticlesHandler")).Host("example.com").Methods("GET").Schemes("http")
	http.ListenAndServe(":8080", r)
}
