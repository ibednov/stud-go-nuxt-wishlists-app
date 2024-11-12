// app/back/services/wishlists/controllers/wishlist_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/models"
)

// GetWishlists возвращает список всех списков подарков
func GetWishlists(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlists []models.Wishlist
		if err := db.Find(&wishlists).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wishlists)
	}
}

// CreateWishlist создает новый список подарков
func CreateWishlist(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlist models.Wishlist
		if err := c.ShouldBindJSON(&wishlist); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&wishlist).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, wishlist)
	}
}
