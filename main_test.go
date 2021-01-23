package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Form a new HTTP request passed to the handler.
	// arg1 = method, arg2 = route, arg3 = request body
	req, err := http.NewRequest("GET", "", nil)

}
