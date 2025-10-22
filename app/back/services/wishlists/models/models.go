// app/back/services/wishlists/models/models.go
package models

import (
	"time" // Импортируем пакет time для работы с временными метками
)

// User представляет модель пользователя
type User struct {
	ID           uint      `gorm:"primaryKey"` // Уникальный идентификатор пользователя
	Username     string    `gorm:"unique;size:255"` // Уникальное имя пользователя
	Email        string    `gorm:"unique;size:255"` // Уникальный адрес электронной почты
	PasswordHash string    `gorm:"size:255"` // Хэш пароля пользователя
	CreatedAt    time.Time // Время создания записи, автоматически управляется GORM
	UpdatedAt    time.Time // Время последнего обновления записи, автоматически управляется GORM
}

// Wishlist представляет модель списка подарков
type Wishlist struct {
	ID          uint      `gorm:"primaryKey"` // Уникальный идентификатор списка подарков
	UserID      uint      `gorm:"index"` // Внешний ключ, указывающий на пользователя
	Name        string    `gorm:"size:255"` // Название списка подарков
	Description *string   // Описание списка подарков, может быть NULL
	CreatedAt   time.Time // Время создания записи, автоматически управляется GORM
	UpdatedAt   time.Time // Время последнего обновления записи, автоматически управляется GORM
}

// Wish представляет модель желания
type Wish struct {
	ID          uint      `gorm:"primaryKey"` // Уникальный идентификатор желания
	WishlistID  uint      `gorm:"index"` // Внешний ключ, указывающий на список подарков
	Name        string    `gorm:"size:255"` // Название желания
	Description *string   // Описание желания, может быть NULL
	Price       float64   `gorm:"type:decimal(10,2)"` // Цена желания
	SourceURL   *string   // URL источника, может быть NULL
	ImageURL    *string   // URL изображения, может быть NULL
	CreatedAt   time.Time // Время создания записи, автоматически управляется GORM
	UpdatedAt   time.Time // Время последнего обновления записи, автоматически управляется GORM
	Status      string    `gorm:"type:enum('available', 'purchased', 'archived')"` // Статус желания
}

// WishlistCollaborator представляет модель коллаборатора списка подарков
type WishlistCollaborator struct {
	ID           uint      `gorm:"primaryKey"` // Уникальный идентификатор коллаборатора
	WishlistID   uint      `gorm:"index"` // Внешний ключ, указывающий на список подарков
	UserID       uint      `gorm:"index"` // Внешний ключ, указывающий на пользователя
	Role         string    `gorm:"type:enum('owner', 'editor', 'viewer')"` // Роль коллаборатора
	Permissions  *string   // Права доступа, может быть NULL
}

// WishlistChange представляет модель изменений в списке подарков
type WishlistChange struct {
	ID                 uint      `gorm:"primaryKey"` // Уникальный идентификатор изменения
	WishlistID        uint      `gorm:"index"` // Внешний ключ, указывающий на список подарков
	ChangeDescription  *string   // Описание изменения, может быть NULL
	ChangedAt         time.Time // Время изменения, автоматически управляется GORM
	UserID            uint      `gorm:"index"` // Внешний ключ, указывающий на пользователя
}

// Tag представляет модель тега
type Tag struct {
	ID   uint   `gorm:"primaryKey"` // Уникальный идентификатор тега
	Name string `gorm:"unique;size:255"` // Название тега
}

// WishlistTag представляет модель связи между списком подарков и тегами
type WishlistTag struct {
	ID         uint `gorm:"primaryKey"` // Уникальный идентификатор связи
	WishlistID uint `gorm:"index"` // Внешний ключ, указывающий на список подарков
	TagID      uint `gorm:"index"` // Внешний ключ, указывающий на тег
}
