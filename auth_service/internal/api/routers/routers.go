package routers

import (
	handlers2 "auth_service/internal/api/handlers"
	"auth_service/internal/api/middlewares"
	"auth_service/internal/database/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, repos *repositories.Repositories) {
	router.GET("/", handlers2.HandleRoot)

	// Adds routes protected by middleware
	router.GET("/protected/resource", middlewares.JWTMiddleware(), handlers2.ProtectedResource)

	router.POST("/accounts/signup", handlers2.SignUp(repos.UserRepo))
	router.POST("/accounts/login", handlers2.SignIn(repos.UserRepo))

	// Add others routers
}
