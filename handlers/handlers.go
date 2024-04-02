// Package handlers contain HTTP handlers to response on requests
// and tell what db method to use
package handlers

import (
	"net/http"
)

// AddHandler func used to add the new record in db
func AddHandler(w http.ResponseWriter, r *http.Request) {

}

// DeleteHandler func used to delete the record in db using an ID
func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

// UpdateHandler func updates the record about book in db using an ID
func UpdateHandler(w http.ResponseWriter, r *http.Request) {

}

// GetByIDHandler func stands for reading data about book with concrete ID
func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

}

// GetHandler func gets the full list of books that are in db
func GetHandler(w http.ResponseWriter, r *http.Request) {

}
