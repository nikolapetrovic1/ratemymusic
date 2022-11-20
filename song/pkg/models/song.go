package models

import (
	"gorm.io/gorm"
	"time"
)

type Song struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Duration    int            `json:"duration"`
	MusicianID  int64          `json:"musician_id"`
	Description string         `json:"description"`
	AlbumID     int64          `json:"album_id"`
	Genres      []Genre        `json:"genres" gorm:"many2many:song_genres;"`
	Audio       string         `json:"audio"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
