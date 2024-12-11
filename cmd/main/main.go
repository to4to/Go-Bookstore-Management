package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/to4to/Go-Bookstore-Management/pkg/routes"
)

func main() {

	r := mux.NewRouter()

	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9000", r))

}
