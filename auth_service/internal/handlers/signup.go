package handlers

import (
	"auth_service/internal/database/repositories"
	"auth_service/pkg/models"
	"auth_service/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
			return
		}

		// Password match check and password confirmation
		if newUser.Password != newUser.ConfirmedPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password do not match"})
			return
		}

		// Checking for existing user by username
		if err := utils.CheckExistingUsername(userRepo, newUser.Username, c); err != nil {
			return
		}

		// Checking for existing user by email
		if err := utils.CheckExistingEmail(userRepo, newUser.Email, c); err != nil {
			return
		}

		// Create the user if the username and email are unique
		if err := userRepo.CreateUser(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}
