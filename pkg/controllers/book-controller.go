package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/to4to/Go-Bookstore-Management/pkg/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBook handles the HTTP request to retrieve a book by its ID.
// It extracts the book ID from the request, fetches the corresponding book
// from the database, and writes the book details to the response.
// If the book is not found, it returns an appropriate error message.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}

func DeleteBook(w http.ResponseWriter, r *http.Request) {}
