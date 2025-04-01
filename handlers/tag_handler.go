package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddTag
// @Tags tag
// @Accept json
// @Produce json
// @Param tag body models.Tag true "Tag object"
// @Success 201 {object} models.Tag
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tag/add [post]
func AddTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
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
// @Success 201 {object} models.Tag
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tag/questions/all [get]
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
