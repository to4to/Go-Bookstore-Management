package config

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestConnect(t *testing.T) {
	// Save the original db value and restore it after the test
	originalDb := db
	defer func() {
		db = originalDb
	}()

	// Test successful connection
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Connect() panicked: %v", r)
			}
		}()

		Connect()
		if db == nil {
			t.Error("Expected db to be initialized, but it is nil")
		}
	}()

	// Test connection failure
	func() {
		// Temporarily simulate a connection failure by using a mock function
		originalDb := db
		defer func() {
			db = originalDb
		}()

		db = nil

		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Connect() to panic, but it did not")
			}
		}()

		Connect()
	}()
}
func TestGetDb(t *testing.T) {
	// Save the original db value and restore it after the test
	originalDb := db
	defer func() {
		db = originalDb
	}()

	// Test when db is initialized
	func() {
		Connect()
		if got := GetDb(); got == nil {
			t.Error("Expected GetDb() to return a non-nil value, but it returned nil")
		}
	}()

	// Test when db is not initialized
	func() {
		db = nil
		if got := GetDb(); got != nil {
			t.Error("Expected GetDb() to return nil, but it returned a non-nil value")
		}
	}()
}
