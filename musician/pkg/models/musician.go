package models

import (
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
	"gorm.io/gorm"
	"time"
)

type Musician struct {
	ID           int64          `json:"id" gorm:"primaryKey"`
	MusicianName string         `json:"musician_name"`
	Name         string         `json:"name"`
	Surname      string         `json:"surname"`
	Songs        []models.Song  `json:"songs" gorm:"foreignKey:MusicianID"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
