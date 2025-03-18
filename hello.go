package main

import (
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var mockQuestions []models.Question

func main() {
	router := gin.Default()
	router.GET("/questions/:id", fetchQuestionById)
	router.GET("/questions", fetchQuestions)
}

func fetchQuestionById(c *gin.Context) {
	id := c.Param("id")
	for _, question := range mockQuestions {
		if id == strconv.FormatInt(question.QuestionId, 10) {
			c.IndentedJSON(http.StatusOK, question)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no question found"})
}

func fetchQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockQuestions)
}
