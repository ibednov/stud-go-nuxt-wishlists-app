package main

import (
	"log" // Импортируем пакет для логирования
	"os"  // Импортируем пакет для работы с окружением

	"wishlists/models" // Импортируем модели из пакета wishlists
	"wishlists/routes" // Импортируем маршруты из пакета wishlists

	"github.com/gin-gonic/gin" // Импортируем фреймворк Gin для создания веб-приложений
	"github.com/joho/godotenv" // Импортируем пакет для загрузки переменных окружения из .env файла
	"gorm.io/driver/mysql"     // Импортируем драйвер MySQL для GORM
	"gorm.io/gorm"             // Импортируем GORM, ORM для Go
)

func init() {
	// Инициализация: загружаем переменные окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file") // Если не удалось загрузить, выводим ошибку и завершаем программу
	}
}

func main() {
	// Получаем параметры подключения к базе данных из переменных окружения
	dbUser := os.Getenv("DB_USER")         // Имя пользователя базы данных
	dbPassword := os.Getenv("DB_PASSWORD") // Пароль пользователя базы данных
	dbName := os.Getenv("DB_NAME")         // Имя базы данных
	dbHost := os.Getenv("DB_HOST")         // Хост базы данных
	dbPort := os.Getenv("DB_PORT")         // Порт базы данных

	// Формируем строку подключения к базе данных
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Открываем соединение с базой данных
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err) // Если не удалось подключиться, выводим ошибку и завершаем программу
	}

	// Автоматическая миграция моделей в базе данных
	if err := db.AutoMigrate(&models.User{}, &models.Wishlist{}, &models.Wish{}, &models.WishlistCollaborator{}, &models.WishlistChange{}, &models.Tag{}, &models.WishlistTag{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err) // Если миграция не удалась, выводим ошибку и завершаем программу
	}

	router := gin.Default() // Создаем новый экземпляр Gin

	// Настройка маршрутов
	routes.SetupRoutes(router, db) // Передаем маршрутизатор и базу данных в функцию настройки маршрутов

	router.Run(":8080") // Запускаем сервер на порту 8080
}
