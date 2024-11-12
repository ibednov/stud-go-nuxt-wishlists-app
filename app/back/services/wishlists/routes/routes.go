// app/back/services/wishlists/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"wishlists/controllers"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/wishlists/users", controllers.GetUsers(db))
	router.GET("/wishlists", controllers.GetWishlists(db))
	router.POST("/wishlists", controllers.CreateWishlist(db))

	router.GET("/wishlists/wishes", controllers.GetWishes(db))
	router.POST("/wishlists/wishes", controllers.CreateWish(db))

	router.GET("/wishlists/collaborators", controllers.GetCollaborators(db))
	router.POST("/wishlists/collaborators", controllers.CreateCollaborator(db))

	router.GET("/wishlists/changes", controllers.GetChanges(db))
	router.POST("/wishlists/changes", controllers.CreateChange(db))

	router.GET("/wishlists/tags", controllers.GetTags(db))
	router.POST("/wishlists/tags", controllers.CreateTag(db))

	router.GET("/wishlists/wishlist_tags", controllers.GetWishlistTags(db))
	router.POST("/wishlists/wishlist_tags", controllers.CreateWishlistTag(db))
}
