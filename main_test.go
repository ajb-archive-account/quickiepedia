package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Form a new HTTP request passed to the handler.
	// arg1 = method, arg2 = route, arg3 = request body
	req, err := http.NewRequest("GET", "", nil)

	// In case of error formatting request, fail & stop the test
	if err != nil {
		t.Fatal(err)
	}

	// HTTP recorder to act as target for request
	recorder := httptest.NewRecorder()

	// Create HTTP handler
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to the handler
	hf.ServeHTTP(recorder, req)

	// Check status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check response is expected
	expected := "Hello, World!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestRouter(t *testing.T) {
	// Instantiate the router using the constructor defined in main.go
	r := newRouter()

	// Create new server using httptest library `NewServer` method
	mockServer := httptest.NewServer(r)

	// GET request to page "/hello"
	resp, err := http.Get(mockServer.URL + "/hello")

	// Handle errors
	if err != nil {
		t.Fatal(err)
	}

	// Log status code 200(ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status %d (OK)", resp.StatusCode)
	}

	// Read response body and convert to string
	defer resp.Body.Close()

	// 1. Read response body into bytes
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Convert bytes to string
	respString := string(b)
	expected := "Hello, World!"

	// Check response matchs handler.
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	// Request to a route we know we didn't define, like the `POST /hello` route.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Check for status 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// Test for an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
