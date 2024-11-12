package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;size:255"`
	Email        string    `gorm:"unique;size:255"`
	PasswordHash string    `gorm:"size:255"`
	CreatedAt    time.Time // Автоматически управляется GORM
	UpdatedAt    time.Time // Автоматически управляется GORM
}

// Wishlist представляет модель списка подарков
type Wishlist struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"index"` // Внешний ключ
	Name        string    `gorm:"size:255"`
	Description *string   // Указатель на строку, позволяющий значениям NULL
	CreatedAt   time.Time // Автоматически управляется GORM
	UpdatedAt   time.Time // Автоматически управляется GORM
}

// Wish представляет модель желания
type Wish struct {
	ID          uint      `gorm:"primaryKey"`
	WishlistID  uint      `gorm:"index"` // Внешний ключ
	Name        string    `gorm:"size:255"`
	Description *string   // Указатель на строку, позволяющий значениям NULL
	Price       float64   `gorm:"type:decimal(10,2)"`
	SourceURL   *string   // Указатель на строку, позволяющий значениям NULL
	ImageURL    *string   // Указатель на строку, позволяющий значениям NULL
	CreatedAt   time.Time // Автоматически управляется GORM
	UpdatedAt   time.Time // Автоматически управляется GORM
	Status      string    `gorm:"type:enum('available', 'purchased', 'archived')"`
}

// WishlistCollaborator представляет модель коллаборатора списка подарков
type WishlistCollaborator struct {
	ID           uint      `gorm:"primaryKey"`
	WishlistID   uint      `gorm:"index"` // Внешний ключ
	UserID       uint      `gorm:"index"` // Внешний ключ
	Role         string    `gorm:"type:enum('owner', 'editor', 'viewer')"`
	Permissions  *string   // Указатель на строку, позволяющий значениям NULL
}

// WishlistChange представляет модель изменений в списке подарков
type WishlistChange struct {
	ID                 uint      `gorm:"primaryKey"`
	WishlistID        uint      `gorm:"index"` // Внешний ключ
	ChangeDescription  *string   // Указатель на строку, позволяющий значениям NULL
	ChangedAt         time.Time // Автоматически управляется GORM
	UserID            uint      `gorm:"index"` // Внешний ключ
}

// Tag представляет модель тега
type Tag struct {
	ID   uint      `gorm:"primaryKey"`
	Name *string   // Указатель на строку, позволяющий значениям NULL
}

// WishlistTag представляет модель связи между списком подарков и тегами
type WishlistTag struct {
	ID         uint `gorm:"primaryKey"`
	WishlistID uint `gorm:"index"` // Внешний ключ
	TagID      uint `gorm:"index"` // Внешний ключ
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

	// Автоматическая миграция моделей
	if err := db.AutoMigrate(&User{}, &Wishlist{}, &Wish{}, &WishlistCollaborator{}, &WishlistChange{}, &Tag{}, &WishlistTag{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := gin.Default()

	// Эндпоинты для работы с пользователями
	router.GET("/wishlists/users", func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	// Эндпоинты для работы со списками подарков
	router.GET("/wishlists", func(c *gin.Context) {
		var wishlists []Wishlist
		if err := db.Find(&wishlists).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wishlists)
	})

	// Эндпоинты для работы с желаниями
	router.GET("/wishlists/wishes", func(c *gin.Context) {
		var wishes []Wish
		if err := db.Find(&wishes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wishes)
	})

	// Эндпоинты для работы с коллабораторами
	router.GET("/wishlists/collaborators", func(c *gin.Context) {
		var collaborators []WishlistCollaborator
		if err := db.Find(&collaborators).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, collaborators)
	})

	// Эндпоинты для работы с изменениями в списках подарков
	router.GET("/wishlists/changes", func(c *gin.Context) {
		var changes []WishlistChange
		if err := db.Find(&changes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, changes)
	})

	// Эндпоинты для работы с тегами
	router.GET("/wishlists/tags", func(c *gin.Context) {
		var tags []Tag
		if err := db.Find(&tags).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tags)
	})

	// Эндпоинты для работы с связями между списками подарков и тегами
	router.GET("/wishlists/wishlist_tags", func(c *gin.Context) {
		var wishlistTags []WishlistTag
		if err := db.Find(&wishlistTags).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, wishlistTags)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/env", func(c *gin.Context) {
		c.String(200, "Hello %s", os.Getenv("NAME"))
	})

	router.Run(":8080")
}
