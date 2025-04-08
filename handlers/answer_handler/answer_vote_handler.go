package answer_handler

import (
	"Learning/database"
	"Learning/helper"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// VoteUpAnswer
// @Tags answer_handler
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Answer
// @Router /api/answer_handler/voteUp/{id} [get]
func VoteUpAnswer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Question
	result := database.DB.First(&answer, id)
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
		answer.Votes += 1
		if updateResult := database.DB.Save(&answer).Error; updateResult != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to vote up"})
			return
		}
		c.IndentedJSON(http.StatusCreated, answer)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not enough reputation to vote"})
	}
}

// VoteDownAnswer
// @Tags answer_handler
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Answer
// @Router /api/answer_handler/voteDown/{id} [get]
func VoteDownAnswer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Question
	result := database.DB.First(&answer, id)
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
		answer.Votes -= 1
		if updateResult := database.DB.Save(&answer).Error; updateResult != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to vote up"})
			return
		}
		c.IndentedJSON(http.StatusCreated, answer)
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not enough reputation to vote"})
	}
}
