package handlers

import (
	"Learning/database"
	"Learning/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddUser
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "Tag object"
// @Success 201 {object} models.User
// @Router /admin/user/add [post]
func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

// DeleteUser
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "Tag object"
// @Success 201 {object} models.User
// @Router /admin/user/delete [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}

	result := database.DB.Delete(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user", "error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, user)

}
