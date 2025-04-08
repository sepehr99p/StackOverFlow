package handlers

import (
	"Learning/database"
	"Learning/helper"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// VoteUpQuestion
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Question
// @Router /api/questions/voteUp/{id} [get]
func VoteUpQuestion(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var question models.Question
	result := database.DB.First(&question, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found"})
		return
	}
	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}
	if user.Reputation > 30 {
		question.Votes += 1
		if updateResult := database.DB.Save(&question).Error; updateResult != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to vote up"})
			return
		}
		c.IndentedJSON(http.StatusCreated, question)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not enough reputation"})
	}
}

// VoteDownQuestion
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Question
// @Router /api/questions/voteDown/{id} [get]
func VoteDownQuestion(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var question models.Question
	result := database.DB.First(&question, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found"})
		return
	}
	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}
	if user.Reputation > 30 {
		question.Votes -= 1
		if updateResult := database.DB.Save(&question).Error; updateResult != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to vote up"})
			return
		}
		c.IndentedJSON(http.StatusCreated, question)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not enough reputation"})
	}
}

// FetchQuestionById
// @Tags questions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Question
// @Router /api/questions/my/{id} [get]
func FetchQuestionById(c *gin.Context) {
	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var question models.Question
	if err := database.DB.First(&question, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Question not found"})
		return
	}

	var answers []models.Answer
	database.DB.Where("question_id = ?", id).Find(&answers)

	var comments []models.Comment
	database.DB.Where("parent_id = ? AND parent_type = ?", id, "question").Find(&comments)

	response := gin.H{
		"user":     user.UserName,
		"question": question,
		"answers":  answers,
		"comments": comments,
	}

	c.IndentedJSON(http.StatusOK, response)
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

	var questionResponses = database.FetchQuestionsWithAnswersAndComments(questions)
	c.IndentedJSON(http.StatusOK, questionResponses)
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
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
