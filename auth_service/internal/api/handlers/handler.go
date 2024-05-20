package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to port :8081!",
	})
}

func ProtectedResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a protected resource"})
}
