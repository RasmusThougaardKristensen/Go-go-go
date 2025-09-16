package main

import (
	"GoSupplyChain/internal/Application/Repositories"
	"GoSupplyChain/internal/Endpoints"
	"GoSupplyChain/internal/Infrastructure/Repositories"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	router := gin.Default()
	var database, err = connectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	userRepository := Repositories.PostgresUserRepository(database)

	userEndpoints(router, userRepository)
	router.Run("localhost:8080")
}

func userEndpoints(router *gin.Engine, repository IRepositories.IUserRepository) {
	getUserEndpoints := Endpoints.GetUserEndpoints(repository)

	router.GET("/users", Endpoints.GetUsersEndpoint)
	router.GET("/users/:id", getUserEndpoints.GetUserEndpoint)
	router.POST("/users", Endpoints.CreateUserEndpoint)
	router.DELETE("/users/:id", Endpoints.DeleteUserEndpoint)
	router.PUT("/users/:id", Endpoints.UpdateUserEndpoint)
}

func connectDatabase() (*sqlx.DB, error) {
	database, err := Repositories.Connect()
	return database, err
}
