// app/back/services/wishlists/controllers/wish_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/models"
)

// GetWishes возвращает список всех желаний
func GetWishes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishes []models.Wish
		if err := db.Find(&wishes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wishes)
	}
}

// CreateWish создает новое желание
func CreateWish(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wish models.Wish
		if err := c.ShouldBindJSON(&wish); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&wish).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, wish)
	}
}
