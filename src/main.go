package main

import (
	"fmt"
	"project-api/config"
	"project-api/controllers"
	"project-api/repositories"
	"project-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	envs := config.GetEnvs()
	db := config.NewConnDB(envs)

	repo := repositories.NewRepository(db)
	services := services.NewService(repo)
	controller := controllers.NewController(services)

	router := gin.New()

	routes := router.Group("/v1")
	routes.GET("/file", controller.Get)

	fmt.Println(envs)
}
