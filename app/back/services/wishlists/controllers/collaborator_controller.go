// app/back/services/wishlists/controllers/collaborator_controller.go
package controllers

import (
	"net/http" // Импортируем пакет для работы с HTTP

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
	"wishlists/models"         // Импортируем модели из пакета wishlists
)

// GetCollaborators возвращает список всех коллаборатора
func GetCollaborators(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var collaborators []models.WishlistCollaborator // Объявляем переменную для хранения списка коллаборатора
		if err := db.Find(&collaborators).Error; err != nil { // Получаем всех коллаборатора из базы данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось получить данные
			return
		}
		c.JSON(http.StatusOK, collaborators) // Возвращаем список коллаборатора с кодом 200
	}
}

// CreateCollaborator создает нового коллаборатора
func CreateCollaborator(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var collaborator models.WishlistCollaborator // Объявляем переменную для нового коллаборатора
		if err := c.ShouldBindJSON(&collaborator); err != nil { // Привязываем JSON из запроса к структуре коллаборатора
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Возвращаем ошибку, если данные некорректны
			return
		}
		if err := db.Create(&collaborator).Error; err != nil { // Сохраняем нового коллаборатора в базе данных
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Возвращаем ошибку, если не удалось сохранить данные
			return
		}
		c.JSON(http.StatusCreated, collaborator) // Возвращаем созданного коллаборатора с кодом 201
	}
}
