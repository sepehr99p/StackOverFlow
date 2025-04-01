package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags answer
// @Accept json
// @Produce json
// @Param answer body models.Answer true "Answer object"
// @Success 201 {object} models.Answer
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /answer/add [post]
func AddAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, answer.UserId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
		return
	}

	var question models.Question
	if err := database.DB.First(&question, answer.QuestionId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Question does not exist"})
		return
	}

	result := database.DB.Create(&answer)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating answer"})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}
