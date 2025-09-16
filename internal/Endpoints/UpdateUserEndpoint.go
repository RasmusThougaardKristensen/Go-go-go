package Endpoints

import (
	"GoSupplyChain/internal/Domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserEndpoint(c *gin.Context) {
	id := c.Param("id")

	var updatedUser models.UserModel

	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i, user := range models.UserModels {
		if user.Id == id {
			updatedUser.Id = id
			models.UserModels[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}
}
