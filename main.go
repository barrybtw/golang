package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Year     string `json:"year"`
}

var books = []Book{
	{ID: "1", Title: "Book 1", Author: "Author 1", Quantity: 1, Year: "2020"},
	{ID: "2", Title: "Book 2", Author: "Author 2", Quantity: 2, Year: "2020"},
	{ID: "3", Title: "Book 3", Author: "Author 3", Quantity: 3, Year: "2020"},
	{ID: "4", Title: "Book 4", Author: "Author 4", Quantity: 4, Year: "2020"},
	{ID: "5", Title: "Book 5", Author: "Author 5", Quantity: 5, Year: "2020"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func createBook(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil { return }
	books = append(books, book)
	c.IndentedJSON(http.StatusCreated, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.Run("localhost:9080")
}