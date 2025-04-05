package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CorrectAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Success 201 {object} models.Answer
// @Router /answer/correctAnswer/{id} [get]
func CorrectAnswer(c *gin.Context) {
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

func FetchAnswersForQuestion(questionId string) []gin.H {
	var answers []models.Answer
	database.DB.Where("question_id = ?", questionId).Find(&answers)
	var answersWithComments []gin.H
	for _, answer := range answers {
		var answerComments []models.Comment
		database.DB.Where("parent_id = ? AND parent_type = ?", answer.AnswerId, "answer").Find(&answerComments)

		answerResponse := gin.H{
			"answer":   answer,
			"comments": answerComments,
		}
		answersWithComments = append(answersWithComments, answerResponse)
	}
	return answersWithComments
}
