package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quanity"`
}

var books = []book{
	{
		ID:       "01",
		Title:    "The Girl with the Dragon Tattoo",
		Author:   "Stieg Larsson",
		Quantity: 81,
	},
	{
		ID:       "02",
		Title:    "The Girl with the Dragon Tattoo",
		Author:   "Stieg Larsson",
		Quantity: 81,
	},
	{
		ID:       "03",
		Title:    "The Girl with the Dragon Tattoo",
		Author:   "Stieg Larsson",
		Quantity: 81,
	},
	{
		ID:       "04",
		Title:    "The Girl with the Dragon Tattoo",
		Author:   "Stieg Larsson",
		Quantity: 81,
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8000")

}
