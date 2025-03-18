package main

import (
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var mockQuestions = []models.Question{}

func main() {
	router := gin.Default()

	router.GET("/questions", fetchQuestions)
}

func fetchQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockQuestions)
}
