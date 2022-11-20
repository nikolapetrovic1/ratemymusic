package models

import (
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	UserId    int64          `json:"user_id"`
	Kind      string         `json:"kind"`
	KindId    int64          `json:"kind_id"`
	Name      string         `json:"name"`
	Favorite  bool           `json:"favorite"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
