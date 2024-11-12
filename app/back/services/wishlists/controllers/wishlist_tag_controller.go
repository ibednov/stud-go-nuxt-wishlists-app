// app/back/services/wishlists/controllers/wishlist_tag_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/models"
)

// GetWishlistTags возвращает список всех связей между списками подарков и тегами
func GetWishlistTags(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlistTags []models.WishlistTag
		if err := db.Find(&wishlistTags).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wishlistTags)
	}
}

// CreateWishlistTag создает новую связь между списком подарков и тегом
func CreateWishlistTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlistTag models.WishlistTag
		if err := c.ShouldBindJSON(&wishlistTag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&wishlistTag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, wishlistTag)
	}
}
