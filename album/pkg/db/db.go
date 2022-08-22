package db

import (
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/models"
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

	db.AutoMigrate(&models.Album{}, models.Genre{})
	db.Create(&models.Genre{Name: "Rock"})
	db.Create(&models.Genre{Name: "Hip-Hop"})
	return Handler{db}
}

func (handler *Handler) SearchAlbums(albumName string) []models.Album {
	var albums []models.Album
	handler.DB.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(albumName)+"%").Find(&albums)
	return albums
}
