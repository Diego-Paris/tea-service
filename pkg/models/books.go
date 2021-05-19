package models

import "github.com/kamva/mgm/v3"

// Book represents each product
type Book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

// CollectionName returns the name of the collection in the database
func (model *Book) CollectionName() string {
	return "my_books"
}

// NewBook takes in the name and number pages and returns an address of a book
func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}
