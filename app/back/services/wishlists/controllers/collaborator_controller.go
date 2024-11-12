// app/back/services/wishlists/controllers/collaborator_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/models"
)

// GetCollaborators возвращает список всех коллаборатора
func GetCollaborators(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var collaborators []models.WishlistCollaborator
		if err := db.Find(&collaborators).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, collaborators)
	}
}

// CreateCollaborator создает нового коллаборатора
func CreateCollaborator(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var collaborator models.WishlistCollaborator
		if err := c.ShouldBindJSON(&collaborator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&collaborator).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, collaborator)
	}
}
