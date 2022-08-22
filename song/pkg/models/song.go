package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Duration    int    `json:"duration"`
	MusicianID  int64  `json:"musician_id"`
	Description string `json:"description"`
	// Features  []int          `json:""`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	AlbumID   int64          `json:"album_id"`
	Genres    []Genre        `json:"genres" gorm:"many2many:song_genres;"`
}
