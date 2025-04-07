package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

// DeleteAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Answer
// @Router /api/answer/delete [delete]
func DeleteAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}
	//todo : check if user has permission to delete answer

	result := database.DB.Delete(&answer)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating answer"})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}

// CorrectAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Answer
// @Router /api/answer/correctAnswer/{id} [get]
func CorrectAnswer(c *gin.Context) {
	//todo : check if user has asked the question to mark it as correct
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Answer
	if err := database.DB.Where("answer_id = ?", id).First(&answer).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "answer not found"})
		return
	}
	answer.IsCorrectAnswer = true
	if updateError := database.DB.Save(&answer).Error; updateError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to update the answer"})
		return
	}
	c.IndentedJSON(http.StatusOK, answer)
}

// AddAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param answer body models.Answer true "Answer object"
// @Success 201 {object} models.Answer
// @Router /api/answer/add [post]
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

	// regex can be updated
	matchString, err := regexp.MatchString("^[]0-9a-zA-Z,!^`@{}=().;/~_|[-]+$", answer.Description)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error checking description"})
		return
	}
	if matchString == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Description containing bad characters"})
		return
	}

	result := database.DB.Create(&answer)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating answer"})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}

// VoteUpAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Answer
// @Router /api/answer/voteUp/{id} [get]
func VoteUpAnswer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var answer models.Question
	result := database.DB.First(&answer, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found"})
		return
	}
	//todo : check if user has enough reputation to vote up a answer
	answer.Votes += 1
	if updateResult := database.DB.Save(&answer).Error; updateResult != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to vote up"})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}
