// app/back/services/wishlists/controllers/wish_controller.go
package controllers

import (
	"net/http" // Импортируем пакет для работы с HTTP

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
	"wishlists/models"         // Импортируем модели из пакета wishlists
)

// GetWishes возвращает список всех желаний
func GetWishes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishes []models.Wish // Объявляем переменную для хранения списка желаний
		if err := db.Find(&wishes).Error; err != nil { // Получаем все желания из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось получить данные
			return
		}
		c.JSON(http.StatusOK, wishes) // Возвращаем список желаний с кодом 200
	}
}

// CreateWish создает новое желание
func CreateWish(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wish models.Wish // Объявляем переменную для нового желания
		if err := c.ShouldBindJSON(&wish); err != nil { // Привязываем JSON из запроса к структуре желания
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку, если данные некорректны
			return
		}
		if err := db.Create(&wish).Error; err != nil { // Сохраняем новое желание в базе данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось сохранить данные
			return
		}
		c.JSON(http.StatusCreated, wish) // Возвращаем созданное желание с кодом 201
	}
}
