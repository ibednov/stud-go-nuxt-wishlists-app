// app/back/services/wishlists/controllers/change_controller.go
package controllers

import (
	"net/http" // Импортируем пакет для работы с HTTP

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
	"wishlists/models"         // Импортируем модели из пакета wishlists
)

// GetChanges возвращает список всех изменений
func GetChanges(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var changes []models.WishlistChange // Объявляем переменную для хранения списка изменений
		if err := db.Find(&changes).Error; err != nil { // Получаем все изменения из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось получить данные
			return
		}
		c.JSON(http.StatusOK, changes) // Возвращаем список изменений с кодом 200
	}
}

// CreateChange создает новое изменение
func CreateChange(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var change models.WishlistChange // Объявляем переменную для нового изменения
		if err := c.ShouldBindJSON(&change); err != nil { // Привязываем JSON из запроса к структуре изменения
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку, если данные некорректны
			return
		}
		if err := db.Create(&change).Error; err != nil { // Сохраняем новое изменение в базе данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось сохранить данные
			return
		}
		c.JSON(http.StatusCreated, change) // Возвращаем созданное изменение с кодом 201
	}
}
