package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect establishes a connection to the MySQL database using GORM.
// It uses the provided connection string to connect to the database.
// If the connection fails, it panics and stops the execution of the program.
func Connect() { //@check-up
	d, err := gorm.Open("mysql", "to4to:admin@123/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

// GetDb returns a pointer to the gorm.DB instance representing the database connection.
// This function can be used to access the database connection throughout the application.
func GetDb() *gorm.DB {
	return db
}
