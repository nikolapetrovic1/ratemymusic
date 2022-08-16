package db

import (
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.SongRating{}, &models.AlbumRating{})
	return Handler{db}
}
