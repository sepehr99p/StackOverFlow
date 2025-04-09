package answer_handler

import (
	"Learning/database"
	"Learning/helper"
	"Learning/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	err := voteUpAnswerWithOwner(&answer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Transaction failed", "error": err.Error()})
		return
	}
}

func voteUpAnswerWithOwner(answer *models.Answer) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var answerOwner models.User
		if err := tx.First(&answerOwner, answer.UserId).Error; err != nil {
			return fmt.Errorf("failed to fetch answer owner: %w", err)
		}

		if err := tx.Model(&answerOwner).Update("reputation", gorm.Expr("reputation + ?", 10)).Error; err != nil {
			return fmt.Errorf("failed to update reputation: %w", err)
		}

		if err := tx.Model(&answer).Update("votes", gorm.Expr("votes + ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to vote up: %w", err)
		}

		//todo log the action later
		return nil
	})
	return err
}

func voteDownAnswerWithOwner(answer *models.Answer) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var answerOwner models.User
		if err := tx.First(&answerOwner, answer.UserId).Error; err != nil {
			return fmt.Errorf("failed to fetch answer owner: %w", err)
		}

		if answerOwner.Reputation <= 0 {
			return fmt.Errorf("answer owner cannot have negative reputation")
		}

		newReputation := answerOwner.Reputation - 10
		if newReputation < 0 {
			newReputation = 0
		}

		if err := tx.Model(&answerOwner).Update("reputation", newReputation).Error; err != nil {
			return fmt.Errorf("failed to update reputation: %w", err)
		}

		if err := tx.Model(&answer).Update("votes", gorm.Expr("votes - ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to downvote: %w", err)
		}

		return nil
	})
	return err
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

	err := voteDownAnswerWithOwner(&answer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Transaction failed", "error": err.Error()})
		return
	}
}
