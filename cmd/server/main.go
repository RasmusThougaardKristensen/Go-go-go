package main

import (
	"GoSupplyChain/internal/Endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
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
