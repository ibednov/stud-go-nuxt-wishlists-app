// app/back/services/wishlists/controllers/tag_controller.go
package controllers

import (
	"net/http" // Импортируем пакет для работы с HTTP

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
	"wishlists/models"         // Импортируем модели из пакета wishlists
)

// GetTags возвращает список всех тегов
func GetTags(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tags []models.Tag // Объявляем переменную для хранения списка тегов
		if err := db.Find(&tags).Error; err != nil { // Получаем все теги из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось получить данные
			return
		}
		c.JSON(http.StatusOK, tags) // Возвращаем список тегов с кодом 200
	}
}

// CreateTag создает новый тег
func CreateTag(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag models.Tag // Объявляем переменную для нового тега
		if err := c.ShouldBindJSON(&tag); err != nil { // Привязываем JSON из запроса к структуре тега
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку, если данные некорректны
			return
		}
		if err := db.Create(&tag).Error; err != nil { // Сохраняем новый тег в базе данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось сохранить данные
			return
		}
		c.JSON(http.StatusCreated, tag) // Возвращаем созданный тег с кодом 201
	}
}
