package models

import (
	"gorm.io/gorm"
	"time"
)

type Album struct {
	ID         int64          `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name"`
	Songs      []Song         `gorm:"foreignKey:AlbumID"`
	MusicianID int64          `json:"musician_id"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
