package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wishlists/models"
	"wishlists/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Автоматическая миграция моделей
	if err := db.AutoMigrate(&models.User{}, &models.Wishlist{}, &models.Wish{}, &models.WishlistCollaborator{}, &models.WishlistChange{}, &models.Tag{}, &models.WishlistTag{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := gin.Default()

	// Настройка маршрутов
	routes.SetupRoutes(router, db)

	router.Run(":8080")
}
