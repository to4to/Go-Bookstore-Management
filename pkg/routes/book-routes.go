package routes

import (
	"github.com/gorilla/mux"
	"github.com/to4to/Go-Bookstore-Management/pkg/controllers"
)

// RegisterRoutes registers the routes for book-related operations with the provided router.
// It sets up the following routes:
// - POST /book: Create a new book (handled by controllers.CreateBook)
// - GET /book: Get a list of all books (handled by controllers.GetBook)
// - GET /book/{bookId}: Get details of a specific book by its ID (handled by controllers.GetBookById)
// - PUT /book/{bookId}: Update details of a specific book by its ID (handled by controllers.UpdateBook)
// - DELETE /book/bookId: Delete a specific book by its ID (handled by controllers.DeleteBook)
var RegisterRoutes = func(routes *mux.Router) {

	routes.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book", controllers.GetBook).Methods("GET")
	
	
}