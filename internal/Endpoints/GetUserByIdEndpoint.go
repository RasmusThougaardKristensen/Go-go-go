package Endpoints

import (
	"GoSupplyChain/internal/Domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserEndpoint(context *gin.Context) {
	id := context.Param("id")

	for _, user := range models.UserModels {
		if user.Id == id {
			context.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
