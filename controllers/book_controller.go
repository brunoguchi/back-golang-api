package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id          int
	Description string
}

func MockedData() []Book {
	books := []Book{
		{Id: 1, Description: "book 1"},
		{Id: 2, Description: "book 2"},
		{Id: 3, Description: "book 3"},
		{Id: 4, Description: "book 4"},
		{Id: 5, Description: "book 5"},
	}

	return books
}

func GetBooks(c *gin.Context) {
	c.JSON(200, MockedData())
}

func GetBookById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{"error": "Id has to be integer"})
		return
	}

	books := MockedData()

	for _, x := range books {
		if x.Id == id {
			c.JSON(200, x)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book of not found"})
}
