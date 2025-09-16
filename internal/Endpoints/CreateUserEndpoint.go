package Endpoints

import (
	"GoSupplyChain/internal/Domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserEndpoint(context *gin.Context) {
	var newUser models.UserModel

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	models.UserModels = append(models.UserModels, newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}
