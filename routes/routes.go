package routes

import (
	"net/http"

	"github.com/MidoriNoKen/latihan-golang-ai/config"
	"github.com/MidoriNoKen/latihan-golang-ai/controller"
	"github.com/MidoriNoKen/latihan-golang-ai/repository"
	"github.com/MidoriNoKen/latihan-golang-ai/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Simple Ping / Health Check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Register API v1 endpoints (only if database is connected)
	if config.DB != nil {
		userRepo := repository.NewUserRepository(config.DB)
		userService := service.NewUserService(userRepo)
		userController := controller.NewUserController(userService)

		v1 := r.Group("/api/v1")
		{
			v1.POST("/users", userController.Register)
			v1.GET("/users", userController.GetAll)
			v1.GET("/users/:id", userController.GetByID)
		}
	}

	return r
}
