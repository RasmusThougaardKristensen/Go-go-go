package Endpoints

import (
	"GoSupplyChain/internal/Domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersEndpoint(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.UserModels)
}
