package db

import (
	"github.com/nikolapetrovic1/ratemymusic/musician/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Musician{})

	return Handler{db}
}

func (handler *Handler) SearchMusician(musicianName string) []models.Musician {
	var musicians []models.Musician
	handler.DB.Where("LOWER(musician_name) LIKE ?", "%"+strings.ToLower(musicianName)+"%").Find(&musicians)
	return musicians
}
