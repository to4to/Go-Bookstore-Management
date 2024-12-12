package routes

import (
	"github.com/gorilla/mux"
	"github.com/to4to/Go-Bookstore-Management/pkg/controllers"
)

var RegisterRoutes = func(routes *mux.Router) {

	routes.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book", controllers.GetBook).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/bookId", controllers.DeleteBook).Methods("DELETE")

}
