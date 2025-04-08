package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// DeleteComment
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param comment body models.Comment true "Comment object"
// @Success 201 {object} models.Comment
// @Router /api/comment/delete [delete]
func DeleteComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid json format"})
		return
	}
	if result := database.DB.Delete(&comment).Error; result != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete comment"})
		return
	}
	c.IndentedJSON(http.StatusOK, comment)
}

// AddComment
// @Tags comment
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param comment body models.Comment true "Comment object"
// @Success 201 {object} models.Comment
// @Router /api/comment/add [post]
func AddComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, comment.UserId).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
		return
	}

	if comment.ParentType == "question_handler" {
		var question models.Question
		if err := database.DB.First(&question, comment.ParentId).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Question not found"})
			return
		}
	} else if comment.ParentType == "answer_handler" {
		var answer models.Answer
		if err := database.DB.First(&answer, comment.ParentId).Error; err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Answer not found"})
			return
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid parent type"})
		return
	}

	// regex can be updated
	matchString, err := regexp.MatchString("^[]0-9a-zA-Z,!^`@{}=().;/~_|[-]+$", comment.Description)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error checking description"})
		return
	}
	if matchString == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Description containing bad characters"})
		return
	}

	result := database.DB.Create(&comment)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating comment"})
		return
	}

	c.IndentedJSON(http.StatusCreated, comment)
}
