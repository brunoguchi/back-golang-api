package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Livro struct {
	Id        int
	Descricao string
}

func MockarDados() []Livro {
	livros := []Livro{
		{Id: 1, Descricao: "livro 1"},
		{Id: 2, Descricao: "livro 2"},
		{Id: 3, Descricao: "livro 3"},
		{Id: 4, Descricao: "livro 4"},
		{Id: 5, Descricao: "livro 5"},
	}

	return livros
}

func GetBooks(c *gin.Context) {
	c.JSON(200, MockarDados())
}

func GetBookById(c *gin.Context) {
	idParam := c.Param("id")
	intVar, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{"error": "Id has to be integer"})
		return
	}

	livros := MockarDados()

	for _, x := range livros {
		if x.Id == intVar {
			c.JSON(200, x)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book of not found"})
}
