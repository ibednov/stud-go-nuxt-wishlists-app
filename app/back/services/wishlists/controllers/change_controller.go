// app/back/services/wishlists/controllers/change_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/models"
)

// GetChanges возвращает список всех изменений
func GetChanges(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var changes []models.WishlistChange
		if err := db.Find(&changes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, changes)
	}
}

// CreateChange создает новое изменение
func CreateChange(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var change models.WishlistChange
		if err := c.ShouldBindJSON(&change); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&change).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, change)
	}
}
