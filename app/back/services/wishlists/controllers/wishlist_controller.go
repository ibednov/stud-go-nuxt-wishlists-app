// app/back/services/wishlists/controllers/wishlist_controller.go
package controllers

import (
	"net/http"         // Импортируем пакет для работы с HTTP
	"wishlists/models" // Импортируем модели из пакета wishlists

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
)

// GetWishlists возвращает список всех списков подарков
func GetWishlists(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlists []models.Wishlist // Объявляем переменную для хранения списка списков подарков
		if err := db.Find(&wishlists).Error; err != nil { // Получаем все списки подарков из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось получить данные
			return
		}
		c.JSON(http.StatusOK, wishlists) // Возвращаем список списков подарков с кодом 200
	}
}

// CreateWishlist создает новый список подарков
func CreateWishlist(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlist models.Wishlist // Объявляем переменную для нового списка подарков
		if err := c.ShouldBindJSON(&wishlist); err != nil { // Привязываем JSON из запроса к структуре списка подарков
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку, если данные некорректны
			return
		}
		if err := db.Create(&wishlist).Error; err != nil { // Сохраняем новый список подарков в базе данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось сохранить данные
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Wishlist created", "data": wishlist}) // Возвращаем созданный список подарков с кодом 201
	}
}

func UpdateWishlist(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var wishlist models.Wishlist
		if err := c.ShouldBindJSON(&wishlist); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&models.Wishlist{}).Where("id = ?", id).Updates(&wishlist).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, gin.H{"message": "Wishlist updated", "data": wishlist})
	}
}

func DeleteWishlist(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id") // Получаем ID списка подарков из параметров запроса
		if err := db.Delete(&models.Wishlist{}, id).Error; err != nil { // Удаляем список подарков из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось удалить данные
			return
		}
		// c.Status(http.StatusNoContent) // Возвращаем статус 204 No Content
		c.JSON(http.StatusOK, gin.H{"message": "Wishlist deleted"}) // Возвращаем сообщение об успешном удалении
	}
}
