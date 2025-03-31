package handlers

import (
	"Learning/database"
	"Learning/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchQuestionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var question models.Question
	result := database.DB.First(&question, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, question)
}

func FetchQuestions(c *gin.Context) {
	var questions []models.Question
	result := database.DB.Find(&questions)

	if result.Error != nil {
		log.Println("Error fetching questions:", result.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving questions"})
		return
	}

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}

func PostQuestion(c *gin.Context) {
	var question models.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, question.UserId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
		return
	}

	result := database.DB.Create(&question)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating question", "error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, question)
}

func FetchMyQuestions(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var userToFind = models.User{UserId: int(id)}
	var questions []models.Question

	result := database.DB.Model(&userToFind).Where("user_id = ?", id).Find(&questions)

	if result.Error != nil {
		log.Println("Error fetching user questions:", result.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user's questions"})
		return
	}

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found for this user"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}
