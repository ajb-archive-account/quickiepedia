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
	staticFileDirectory := http.Dir("/assets/")

	// `stripPrefix` method to remove the "/assets/" prefix when looking for files.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Router is now formed by calling the `newRouter` constructor defined above.
func main() {
	fmt.Println("Running on :8080/hello\nCtrl+C to close")
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
