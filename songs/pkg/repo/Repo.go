package repo

import (
	"github.com/nikolapetrovic1/ratemymusic/songs/pkg/model"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	log "github.com/go-kit/log"
)

type SongRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func New() (*SongRepository, error) {
	repo := &SongRepository{}

	dsn := "user=postgres password=admin dbname=SONG_DB port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repo.db = db
	repo.db.AutoMigrate(&model.Song{})
	return repo, nil
}

func (repo *SongRepository) Save(song model.Song) (model.Song, error) {
	//result :=repo.db.Create(song)
	repo.db.Create(&song)
	return song, nil
}

func (repo *SongRepository) FindAll() ([]model.Song, error) {
	var songs []model.Song
	repo.db.Find(&songs)
	return songs, nil
}
