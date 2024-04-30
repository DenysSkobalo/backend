package handlers

import (
	"auth_service/internal/database/repositories"
	"auth_service/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to port :8081!",
	})
}

func SignUp(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
			return
		}

		if err := userRepo.CreateUser(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}
