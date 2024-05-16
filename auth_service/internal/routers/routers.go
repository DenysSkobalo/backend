package routers

import (
	"auth_service/internal/database/repositories"
	"auth_service/internal/handlers"
	"auth_service/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, repos *repositories.Repositories) {
	router.GET("/", handlers.HandleRoot)

	// Adds routes protected by middleware
	router.GET("/protected/resource", middlewares.JWTMiddleware(), handlers.ProtectedResource)

	router.POST("/signup", handlers.SignUp(repos.UserRepo))
	router.POST("/login", handlers.Login(repos.UserRepo))

	// Add others routers
}
