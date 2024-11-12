package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Определение модели
type YourModel struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	// Добавьте другие поля по необходимости
}

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

	// Используйте переменную db для выполнения операций с базой данных
	if err := db.AutoMigrate(&YourModel{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/env", func(c *gin.Context) {
		c.String(200, "Hello %s", os.Getenv("NAME"))
	})

	router.Run(":8080")
}
