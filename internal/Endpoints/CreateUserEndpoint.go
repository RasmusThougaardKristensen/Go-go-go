package Endpoints

import (
	IRepositories "GoSupplyChain/internal/Application/Repositories"
	"GoSupplyChain/internal/Domain/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	userRepository IRepositories.IUserRepository
}

func CreateUserEndpoint(userRepository IRepositories.IUserRepository) *UserRepository {
	return &UserRepository{
		userRepository: userRepository,
	}
}

func (userRepository *UserRepository) CreateUserEndpoint(context *gin.Context) {
	var createUserRequest models.CreateUserRequest

	if err := context.BindJSON(&createUserRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
		return
	}

	user := models.User(createUserRequest.Name, createUserRequest.Email)

	err := userRepository.userRepository.CreateUser(&user)

	if err != nil {
		// Handle different types of errors
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			context.JSON(http.StatusConflict, gin.H{
				"error": "User with this email already exists",
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create user",
				"message": err.Error(),
			})
		}
		return
	}

	context.JSON(http.StatusCreated, user)
}
