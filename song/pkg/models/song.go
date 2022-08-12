package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Duration   int    `json:"duration"`
	MusicianID int64  `json:"musician_id"`
	// Features  []int          `json:""`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
