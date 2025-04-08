package question_handler

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
