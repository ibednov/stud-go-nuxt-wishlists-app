// app/back/services/wishlists/routes/routes.go
package routes

import (
	"wishlists/controllers" // Импортируем контроллеры из пакета wishlists

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
)

// SetupRoutes настраивает маршруты для приложения
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/wishlists/users", controllers.GetUsers(db)) // Получение списка пользователей
	router.GET("/wishlists", controllers.GetWishlists(db)) // Получение списка всех списков подарков
	router.POST("/wishlists", controllers.CreateWishlist(db)) // Создание нового списка подарков
	router.DELETE("/wishlists/:id", controllers.DeleteWishlist(db)) // Удаление списка подарков

	router.GET("/wishlists/wishes", controllers.GetWishes(db)) // Получение списка всех желаний
	router.POST("/wishlists/wishes", controllers.CreateWish(db)) // Создание нового желания

	router.GET("/wishlists/collaborators", controllers.GetCollaborators(db)) // Получение списка коллаборатора
	router.POST("/wishlists/collaborators", controllers.CreateCollaborator(db)) // Создание нового коллаборатора

	router.GET("/wishlists/changes", controllers.GetChanges(db)) // Получение списка изменений
	router.POST("/wishlists/changes", controllers.CreateChange(db)) // Создание нового изменения

	router.GET("/wishlists/tags", controllers.GetTags(db)) // Получение списка тегов
	router.POST("/wishlists/tags", controllers.CreateTag(db)) // Создание нового тега

	router.GET("/wishlists/wishlist_tags", controllers.GetWishlistTags(db)) // Получение списка связей между списками подарков и тегами
	router.POST("/wishlists/wishlist_tags", controllers.CreateWishlistTag(db)) // Создание новой связи между списком подарков и тегом
}
