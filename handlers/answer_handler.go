package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, answer.UserId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
		return
	}

	var question models.Question
	if err := database.DB.First(&question, answer.QuestionId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Question does not exist", "error": err.Error()})
		return
	}

	result := database.DB.Create(&answer)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating answer", "error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}
