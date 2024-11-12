// app/back/services/wishlists/controllers/wishlist_tag_controller.go
package controllers

import (
	"net/http" // Импортируем пакет для работы с HTTP

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
	"wishlists/models"         // Импортируем модели из пакета wishlists
)

// GetWishlistTags возвращает список всех связей между списками подарков и тегами
func GetWishlistTags(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlistTags []models.WishlistTag // Объявляем переменную для хранения списка связей
		if err := db.Find(&wishlistTags).Error; err != nil { // Получаем все связи из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось получить данные
			return
		}
		c.JSON(http.StatusOK, wishlistTags) // Возвращаем список связей с кодом 200
	}
}

// CreateWishlistTag создает новую связь между списком подарков и тегом
func CreateWishlistTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlistTag models.WishlistTag // Объявляем переменную для новой связи
		if err := c.ShouldBindJSON(&wishlistTag); err != nil { // Привязываем JSON из запроса к структуре связи
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку, если данные некорректны
			return
		}
		if err := db.Create(&wishlistTag).Error; err != nil { // Сохраняем новую связь в базе данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось сохранить данные
			return
		}
		c.JSON(http.StatusCreated, wishlistTag) // Возвращаем созданную связь с кодом 201
	}
}
