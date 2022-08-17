package models

import (
	"gorm.io/gorm"
	"time"
)

type SongRating struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	UserID      int64          `json:"user_id" gorm:"not null"`
	RatingValue float64        `json:"rating_value" gorm:"not null"`
	SongID      int64          `json:"song_id" gorm:"not null"`
	Review      string         `json:"review"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type AlbumRating struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	UserID      int64          `json:"user_id" gorm:"not null"`
	RatingValue float64        `json:"rating_value" gorm:"not null"`
	AlbumId     int64          `json:"album_id" gorm:"not null"`
	Review      string         `json:"review"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
