package answer_handler

import (
	"Learning/database"
	"Learning/helper"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

// DeleteAnswer
// @Tags answer_handler
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Answer
// @Router /api/answer_handler/delete [delete]
func DeleteAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}
	//todo : check if user has permission to delete answer_handler

	result := database.DB.Delete(&answer)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating answer_handler"})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}

// CorrectAnswer
// @Tags answer_handler
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Answer
// @Router /api/answer_handler/correctAnswer/{id} [get]
func CorrectAnswer(c *gin.Context) {
	//todo : check if user has asked the question_handler to mark it as correct
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Answer
	if err := database.DB.Where("answer_id = ?", id).First(&answer).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "answer_handler not found"})
		return
	}

	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	var question models.Question
	if questionQueryError := database.DB.Where("question_id = ?", answer.QuestionId).First(&question).Error; questionQueryError != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to fetch answer's question"})
		return
	}

	if user.UserId == question.UserId {
		answer.IsCorrectAnswer = true
		if updateError := database.DB.Save(&answer).Error; updateError != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to update the answer_handler"})
			return
		}
		c.IndentedJSON(http.StatusOK, answer)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "only the user asked the question, may mark it as correct"})
	}

}

// AddAnswer
// @Tags answer_handler
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param answer_handler body models.Answer true "Answer object"
// @Success 201 {object} models.Answer
// @Router /api/answer_handler/add [post]
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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating answer_handler"})
		return
	}
	c.IndentedJSON(http.StatusCreated, answer)
}
