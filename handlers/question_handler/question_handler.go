package question_handler

import (
	"Learning/database"
	"Learning/helper"
	"Learning/models"
	"Learning/public"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FetchQuestionById
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Question
// @Router /api/questions/{id} [get]
func FetchQuestionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var question []models.Question
	if err := database.DB.First(&question, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, database.FetchQuestionsWithAnswersAndComments(question))
}

// FetchQuestions
// @Tags questions
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Question
// @Router /api/questions/all [get]
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
	c.IndentedJSON(http.StatusOK, database.FetchQuestionsWithAnswersAndComments(questions))
}

// DeleteQuestion
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param question body models.Question true "Question object"
// @Success 201 {object} models.Question
// @Router /api/questions/delete [delete]
func DeleteQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": public.InvalidJson, "error": err.Error()})
		return
	}

	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	if user.IsAdmin || user.UserId == question.UserId {
		result := database.DB.Delete(&question)
		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting question", "error": result.Error.Error()})
			return
		}
		c.IndentedJSON(http.StatusAccepted, question)
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "you can only delete your own questions"})
	}

}

// PostQuestion
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param question body models.Question true "Question Data"
// @Success 201 {object} models.Question
// @Router /api/questions/add [post]
func PostQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": public.InvalidJson, "error": err.Error()})
		return
	}
	if helper.FetchUserFromToken(c.GetHeader("Authorization")) == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	result := database.DB.Create(&question)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating question", "error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, question)
}

// FetchMyQuestions
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Question
// @Router /api/questions/my [get]
func FetchMyQuestions(c *gin.Context) {
	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}
	var questions []models.Question
	result := database.DB.Model(&user).Find(&questions)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user's questions", "error": result.Error.Error()})
		return
	}

	if len(questions) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No questions found for this user"})
		return
	}

	c.IndentedJSON(http.StatusOK, database.FetchQuestionsWithAnswersAndComments(questions))
}
