package model

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Musician int    `json:"musician_id"`
	// Features  []int          `json:""`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// func (t Todo) String() string {
// 	b, err := json.Marshal(t)
// 	if err != nil {
// 	   return "unsupported value type"
// 	}
// 	return string(b)
//  }
