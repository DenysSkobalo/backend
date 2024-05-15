package routers

import (
	"auth_service/internal/database/repositories"
	"auth_service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, repos *repositories.Repositories) {
	router.GET("/", handlers.HandleRoot)

	router.POST("/signup", handlers.SignUp(repos.UserRepo))
	router.POST("/login", handlers.Login(repos.UserRepo))

	// Add others routers
}
