package Endpoints

import (
	"GoSupplyChain/internal/Application/Repositories"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetUserByIdEndpoint struct {
	userRepository IRepositories.IUserRepository
}

func GetUserEndpoints(userRepo IRepositories.IUserRepository) *GetUserByIdEndpoint {
	return &GetUserByIdEndpoint{
		userRepository: userRepo,
	}
}

func (endpoint *GetUserByIdEndpoint) GetUserEndpoint(context *gin.Context) {
	id := context.Param("id")
	user, err := endpoint.userRepository.GetUserById(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			context.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		} else if strings.Contains(err.Error(), "invalid UUID format") {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID format",
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	context.JSON(http.StatusOK, user)
}
