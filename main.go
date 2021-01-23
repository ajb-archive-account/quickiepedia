package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	fmt.Println("\n	   Running on :8080/hello\n	   Ctrl+C to close\n")

	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
