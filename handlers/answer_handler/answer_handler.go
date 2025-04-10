package answer_handler

import (
	"Learning/database"
	"Learning/database/db_helper"
	"Learning/helper"
	"Learning/models"
	"Learning/public"
	"fmt"
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": public.InvalidJson})
		return
	}

	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	if user.IsAdmin || user.UserId == answer.UserId {
		result := database.DB.Delete(&answer)
		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting answer"})
			return
		}
		c.IndentedJSON(http.StatusCreated, answer)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you can only delete your own answers"})
	}
}

// CorrectAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Answer
// @Router /api/answer/correctAnswer/{id} [get]
func CorrectAnswer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Answer
	if err := database.DB.Where("answer_id = ?", id).First(&answer).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "answer not found"})
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

	if user.UserId != question.UserId {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "only the user asked the question, may mark it as correct"})
		return
	}

	err := db_helper.MarkAnswerAsCorrect(&answer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Transaction failed", "error": err.Error()})
		return
	}
	log := models.Log{
		UserID:           uint(user.UserId),
		Action:           "mark_answer_as_correct",
		EntityType:       "answer",
		EntityID:         uint(answer.AnswerId),
		ReputationChange: 10,
		Description:      fmt.Sprintf("User gained 10 reputation for vote up on answer %d", answer.AnswerId),
	}
	database.SaveLog(&log)
	c.IndentedJSON(http.StatusOK, &answer)

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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": public.InvalidJson})
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
