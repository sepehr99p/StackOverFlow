package answer_handler

import (
	"Learning/database"
	"Learning/database/db_helper"
	"Learning/helper"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	var answer models.Answer
	result := database.DB.First(&answer, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Answer not found"})
		return
	}
	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	if user.Reputation < 30 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not enough reputation to vote"})
		return
	}

	err := db_helper.VoteUpAnswerWithOwner(&answer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Transaction failed", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "answer voted up"})
}

// VoteDownAnswer
// @Tags answer
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param id path string true "id"
// @Success 201 {object} models.Answer
// @Router /api/answer/voteDown/{id} [get]
func VoteDownAnswer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Answer
	result := database.DB.First(&answer, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Answer not found"})
		return
	}
	user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
	if user == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
		return
	}

	if user.Reputation < 30 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not enough reputation to vote"})
		return
	}

	err := db_helper.VoteDownAnswerWithOwner(&answer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Transaction failed", "error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "answer voted down"})
}
