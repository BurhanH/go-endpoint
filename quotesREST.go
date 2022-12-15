package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quote struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

var Quotes []Quote

func returnQuotes(c *gin.Context) {
	fmt.Println("Endpoint Hit: returnAllQuotes")
	c.IndentedJSON(http.StatusOK, Quotes)
}

func returnQuote(c *gin.Context) {
	fmt.Println("Endpoint Hit: returnQuote")
	id := c.Param("id")

	for _, q := range Quotes {
		if q.Id == id {
			c.IndentedJSON(http.StatusOK, q)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Quote not found"})
}

func createQuote(c *gin.Context) {
	fmt.Println("Endpoint Hit: createQuote")
	var newQuote Quote
	if err := c.BindJSON(&newQuote); err != nil {
		return
	}

	Quotes = append(Quotes, newQuote)
	c.IndentedJSON(http.StatusCreated, newQuote)
}

func deleteQuote(c *gin.Context) {
	fmt.Println("Endpoint Hit: deleteQuote")
	id := c.Param("id")

	for i, q := range Quotes {
		if q.Id == id {
			Quotes = append(Quotes[:i], Quotes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "Quote deleted"})
			return
		}
	}
}

func main() {

	Quotes = []Quote{
		{Id: "0", Author: "Dr. Seuss", Quote: "Don't cry because it's over, smile because it happened."},
		{Id: "1", Author: "Oscar Wilde", Quote: "Be yourself; everyone else is already taken."},
		{Id: "2", Author: "Albert Einstein", Quote: "Two things are infinite: the universe and human stupidity; and I'm not sure about the universe."},
	}

	router := gin.Default()
	router.GET("/quotes", returnQuotes)
	router.GET("/quotes/:id", returnQuote)
	router.POST("/quotes", createQuote)
	router.DELETE("/quotes/:id", deleteQuote)

	router.Run("localhost:8090")
}
