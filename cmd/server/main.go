package main

import (
	"GoSupplyChain/internal/Endpoints"
	"GoSupplyChain/internal/Infrastructure/Repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	connectDatabase()
	userEndpoints(router)
	router.Run("localhost:8080")
}

func userEndpoints(router *gin.Engine) {
	router.GET("/users", Endpoints.GetUsersEndpoint)
	router.GET("/users/:id", Endpoints.GetUserEndpoint)
	router.POST("/users", Endpoints.CreateUserEndpoint)
	router.DELETE("/users/:id", Endpoints.DeleteUserEndpoint)
	router.PUT("/users/:id", Endpoints.UpdateUserEndpoint)
}

func connectDatabase() {
	Repositories.Connect()
}
