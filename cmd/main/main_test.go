package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/to4to/Go-Bookstore-Management/pkg/routes"
)

// TestMain is a test function that sets up a new HTTP server using the
// Gorilla Mux router and the application's routes. It sends a GET request
// to the server and checks if the response status code is 200 OK. If the
// request fails or the status code is not 200, the test will fail.
func TestMain(t *testing.T) {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}
}
