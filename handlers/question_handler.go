package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary fetch question by id
// @Description fetch question by id
// @Tags questions
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /questions/my/{id} [get]
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

// @Summary Fetch all questions
// @Description fetch all the questions that have been asked
// @Tags questions
// @Accept json
// @Produce json
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string"
// @Failure 500 {object} map[string]string"
// @Router /questions/all [get]
func FetchQuestions(c *gin.Context) {
	var questions []models.Question
	result := database.DB.Find(&questions)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving questions", "error": result.Error.Error()})
		return
	}

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}

// @Summary Delete a questions
// @Description Deleting the selected questions
// @Tags questions
// @Accept json
// @Produce json
// @Param question body models.Question true "Question object"
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /questions/delete [post]
func DeleteQuestion(c *gin.Context) {
	var question models.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	//todo : check if the user has permission to delete question
	result := database.DB.Delete(&question)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting question", "error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, question)
}

// postQuestion creates a new question
// @Summary Add a new question
// @Description Allows users to post a new question with tags, description, and votes
// @Tags questions
// @Accept json
// @Produce json
// @Param question body models.Question true "Question Data" // "question" is the body content for creating a new question
// @Success 201 {object} models.Question // Success response with the created question
// @Failure 400 {object} map[string]string // Error if bad request
// @Failure 500 {object} map[string]string // Internal server error if something goes wrong
// @Router /questions/add [post]
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

// @Summary Fetch questions asked by my user
// @Description fetch questions asked by my user
// @Tags questions
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 201 {object} models.Question
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /questions/my/{user_id} [get]
func FetchMyQuestions(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var userToFind = models.User{UserId: int(id)}
	var questions []models.Question

	result := database.DB.Model(&userToFind).Where("user_id = ?", id).Find(&questions)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user's questions", "error": result.Error.Error()})
		return
	}

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found for this user"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)
}
