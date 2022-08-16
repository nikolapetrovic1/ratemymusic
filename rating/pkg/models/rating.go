package models

import (
	"gorm.io/gorm"
	"time"
)

type SongRating struct {
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
	RatingValue float64        `json:"rating_value"`
	SongID      int64          `json:"song_id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type AlbumRating struct {
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
	RatingValue float64        `json:"rating_value"`
	AlbumId     int64          `json:"album_id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
