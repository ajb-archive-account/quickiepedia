package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Create and return the router
// split from main() to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Router is now formed by calling the `newRouter` constructor defined above.
func main() {
	fmt.Println("\n	   Running on :8080/hello\n	   Ctrl+C to close\n")
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
