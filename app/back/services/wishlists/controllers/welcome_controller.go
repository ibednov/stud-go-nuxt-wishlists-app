// app/back/services/wishlists/controllers/welcome_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Wishlists API"})
}
