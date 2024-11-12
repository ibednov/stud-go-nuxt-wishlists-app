// app/back/services/wishlists/controllers/tag_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/models"
)

// GetTags возвращает список всех тегов
func GetTags(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tags []models.Tag
		if err := db.Find(&tags).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tags)
	}
}

// CreateTag создает новый тег
func CreateTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag models.Tag
		if err := c.ShouldBindJSON(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&tag).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, tag)
	}
}
