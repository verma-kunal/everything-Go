package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// store books info
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// slice of books
var books = []Book{
	{ID: "1", Title: "Harry Potter: Deathly Hollows P1", Author: "JK Rowling", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 3},
}

// returns all books
func GetBooks(c *gin.Context) {

	// json object + http status code
	c.IndentedJSON(http.StatusOK, books)
}

// get book by ID
func GetBookById(c *gin.Context){
	// set the id param - '/books/{id}'
	id := c.Param("id")
	book, err := BookById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// find book by ID
func BookById(id string) (*Book, error) {

	// looping through the slice
	for i, abook := range books {
		if abook.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// create book
func CreateBook(c *gin.Context) {

	var newBook Book

	// bind the json of request to the 'newBook' var
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// append the json data to the book slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}

func main() {

	router := gin.Default()

	// GET route
	router.GET("/books", GetBooks)

	// GET by ID
	router.GET("/books/:id", GetBookById)

	// POST route
	router.POST("/books", CreateBook)

	// run server
	router.Run("localhost:8080")

}
