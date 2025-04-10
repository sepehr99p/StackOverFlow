package handlers

import (
	"Learning/database"
	"Learning/error"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddTag
// @Tags tag
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param tag body models.Tag true "Tag object"
// @Success 201 {object} models.Tag
// @Router /api/tag/add [post]
func AddTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": error.InvalidJson, "error": err.Error()})
		return
	}

	result := database.DB.Create(&tag)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error creating tag", "error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, tag)
}

// FetchTagQuestions
// @Tags tag
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} models.Tag
// @Router /api/tag/questions/all [get]
func FetchTagQuestions(c *gin.Context) {
	tagName := c.Param("name")
	var questions []models.Question
	result := database.DB.First(&questions, tagName)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error fetching questions"})
		return
	}

	c.IndentedJSON(http.StatusOK, questions)

}
