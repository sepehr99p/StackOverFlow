package question_handler

import (
	"Learning/database"
	"Learning/models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// SearchQuestions
// @Tags question
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param query query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {array} models.Question
// @Router /api/questions/search [get]
func SearchQuestions(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	cacheKey := "search:" + query + ":page:" + pageStr + ":limit:" + limitStr

	ctx := context.Background()
	cachedResult, err := database.GetCachedQuestion(ctx, cacheKey)
	if err == nil {
		var questions []models.Question
		if err := json.Unmarshal([]byte(cachedResult), &questions); err == nil {
			c.JSON(http.StatusOK, questions)
			return
		}
	}

	words := strings.Fields(query)
	var searchTerms []string
	for _, word := range words {
		if len(word) > 2 {
			searchTerms = append(searchTerms, "%"+word+"%")
		}
	}

	var questions []models.Question
	queryBuilder := database.DB.Model(&models.Question{})

	for _, term := range searchTerms {
		queryBuilder = queryBuilder.Where("description LIKE ?", term)
	}

	offset := (page - 1) * limit
	err = queryBuilder.
		Preload("Tags").
		Order("votes DESC").
		Offset(offset).
		Limit(limit).
		Find(&questions).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search questions"})
		return
	}

	if jsonData, err := json.Marshal(questions); err == nil {
		database.CacheQuestion(ctx, cacheKey, string(jsonData))
	}

	c.JSON(http.StatusOK, questions)
}
