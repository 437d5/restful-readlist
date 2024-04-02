package models

// import (
// 	"encoding/json"
// )

type Book struct {
	Done          bool   `json:"done"`
	Author        string `json:"author"`
	Title         string `json:"title"`
	YearPublished string `json:"year_published"`
	Rating        string `json:"rating"`
}
