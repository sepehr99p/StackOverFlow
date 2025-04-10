package handlers

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReportInteraction
// @Tags report
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param report body models.Report true "Report object"
// @Success 201 {object} models.Report
// @Router /api/report [post]
func ReportInteraction(c *gin.Context) {
	var report models.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid json format"})
		return
	}
	if report.ReportType == "comment" {
		var comment models.Comment
		if commentQueryResult := database.DB.First(&comment, report.ReportId); commentQueryResult.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "comment not found"})
			return
		}
	} else if report.ReportType == "answer" {
		var answer models.Answer
		if commentQueryResult := database.DB.First(&answer, report.ReportId); commentQueryResult.Error != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "answer not found"})
			return
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid parent type"})
		return
	}
	if reportCreationResult := database.DB.Create(&report); reportCreationResult.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create report"})
		return
	}
	c.IndentedJSON(http.StatusCreated, &report)
}
