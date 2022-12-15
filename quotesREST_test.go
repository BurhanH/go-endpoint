package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllQuotes(t *testing.T) {
	r := SetUpRouter()
	r.GET("/quotes", returnQuotes)
	req, _ := http.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var quotes []Quote
	json.Unmarshal(w.Body.Bytes(), &quotes)

	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Assert(t, len(quotes) != 0)
}

func TestAddQuote(t *testing.T) {
	r := SetUpRouter()
	r.POST("quotes", createQuote)
	quote := Quote{
		Id:     "10",
		Author: "Test Author",
		Quote:  "Test Quote",
	}
	jsonValue, _ := json.Marshal(quote)
	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
