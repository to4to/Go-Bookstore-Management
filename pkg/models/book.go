package models

import (
	"github.com/jinzhu/gorm"
	"github.com/to4to/Go-Bookstore-Management/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

// init initializes the database connection and performs the automatic migration
// for the Book model. It is called automatically when the package is imported.
func init() {
	config.Connect()
	db := config.GetDb()
	db.AutoMigrate(&Book{})
}

// CreateBook creates a new book record in the database.
// It returns the created Book instance.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records from the database.
// It returns a slice of Book structs containing the details of each book.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById retrieves a book from the database by its ID.
// It takes an integer ID as a parameter and returns a pointer to the Book
// and a pointer to the gorm.DB instance.
// If the book is found, it is stored in the getBook variable and returned.
// If the book is not found, the returned Book pointer will be nil.
func GetBookById(Id int64) (*Book, *gorm.DB) {

	var getBook Book

	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book record from the database based on the provided ID.
// It returns the deleted Book object.
//
// Parameters:
//   - ID: The ID of the book to be deleted.
//
// Returns:
//   - Book: The deleted Book object.
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
