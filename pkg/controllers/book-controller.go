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


	func GetBookById(w http.ResponseWriter, r *http.Request) {
    // Get book ID from URL parameters
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Get book details from model
    book, _ := models.GetBookById(ID)

    // Return JSON response
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) 
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Get book ID from URL parameters
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Parse request body
    var updateBook models.Book
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&updateBook); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Update book in database
    bookDetails, db := models.GetBookById(ID)
    if updateBook.Name != "" {
        bookDetails.Name = updateBook.Name
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publication != "" {
        bookDetails.Publication = updateBook.Publication
    }
    db.Save(&bookDetails)

    // Return updated book
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    // Get book ID from URL parameters
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Delete book from database
    book := models.DeleteBook(ID)
    
    // Return deleted book details
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}	

