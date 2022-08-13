package db

import (
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
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

	db.AutoMigrate(&models.Song{})

	return Handler{db}
}

func (handler *Handler) SearchSongs(songName string) []models.Song {
	var songs []models.Song
	handler.DB.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(songName)+"%").Find(&songs)
	return songs
}
