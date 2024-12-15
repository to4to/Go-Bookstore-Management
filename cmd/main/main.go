package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/to4to/Go-Bookstore-Management/pkg/routes"
)

// main is the entry point of the application. It initializes a new router using the Gorilla Mux package,
// registers the application routes, and starts an HTTP server listening on localhost at port 9000.
func main() {

	r := mux.NewRouter()

	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9000", r))

}
