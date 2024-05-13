package handlers

import (
	"auth_service/internal/database/repositories"
	"auth_service/pkg/models"
	"database/sql"
	"errors"
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

		// Password match check and password confirmation
		if newUser.Password != newUser.ConfirmedPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password do not match"})
			return
		}

		// Checking for existing user by username
		if err := checkExistingUsername(userRepo, newUser.Username, c); err != nil {
			return
		}

		// Checking for existing user by email
		if err := checkExistingEmail(userRepo, newUser.Email, c); err != nil {
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

func checkExistingUsername(userRepo *repositories.UserRepository, username string, c *gin.Context) error {
	existingUser, err := userRepo.GetUserByUsername(username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return err
	}
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return errors.New("username already exists")
	}
	return nil
}

func checkExistingEmail(userRepo *repositories.UserRepository, email string, c *gin.Context) error {
	existingEmailUser, err := userRepo.GetUserByEmail(email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return err
	}
	if existingEmailUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return errors.New("email already exists")
	}
	return nil
}
