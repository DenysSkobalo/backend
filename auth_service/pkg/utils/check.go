package utils

import (
	"auth_service/internal/database/repositories"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckExistingUsername(userRepo repositories.UserRepository, username string, c *gin.Context) error {
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

func CheckExistingEmail(userRepo repositories.UserRepository, email string, c *gin.Context) error {
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
