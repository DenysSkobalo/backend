package handlers

import (
	"auth_service/internal/database/repositories"
	"auth_service/pkg/models"
	"auth_service/pkg/utils"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Login(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser models.User
		if err := c.ShouldBindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
			return
		}

		dbUser, err := userRepo.GetUserByUsername(loginUser.Username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
	}
}
