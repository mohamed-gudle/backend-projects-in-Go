package shorturl

import (
	"time"

	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	ShortCode string `gorm:"unique;not null"`
	OriginalURL string `gorm:"not null"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
