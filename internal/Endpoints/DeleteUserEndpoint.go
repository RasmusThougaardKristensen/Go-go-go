package Endpoints

import (
	"GoSupplyChain/internal/Domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserEndpoint(context *gin.Context) {
	id := context.Param("id")

	for i, user := range models.UserModels {
		if user.Id == id {
			models.UserModels = append(models.UserModels[:i], models.UserModels[i+1:]...)
			context.IndentedJSON(http.StatusOK, user)
			return
		}
	}
}
