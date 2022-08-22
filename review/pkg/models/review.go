package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID        int64 `gorm:"primaryKey"`
	UserId    int64
	Text      string
	Type      string
	TypeId    int64
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
