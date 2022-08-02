package service

import (
	"context"

	"github.com/nikolapetrovic1/ratemymusic/songs/pkg/model"
	"github.com/nikolapetrovic1/ratemymusic/songs/pkg/repo"
)

// SongsService describes the service.
type SongsService interface {
	// Add your methods here
	Get(ctx context.Context) (model.Song, error)
	GetAll(ctx context.Context) ([]model.Song, error)
	Add(ctx context.Context, song model.Song) (model.Song, error)
	Delete(ctx context.Context, id string) error
}
type basicSongsService struct {
	songRepo *repo.SongRepository
}

func (b *basicSongsService) Get(ctx context.Context) (model.Song, error) {
	// TODO implement the business logic of Get
	pesma := model.Song{Name: "Pera"}
	return pesma, nil
}
func (b *basicSongsService) Add(ctx context.Context, song model.Song) (model.Song, error) {
	// TODO implement the business logic of Add
	song, err := b.songRepo.Save(song)
	return song, err
}
func (b *basicSongsService) Delete(ctx context.Context, id string) (e0 error) {
	// TODO implement the business logic of Delete
	return e0
}

// NewBasicSongsService returns a naive, stateless implementation of SongsService.
func NewBasicSongsService() SongsService {
	songRepo, _ := repo.New()
	songRepo.Save(model.Song{Name: "Feel like god"})
	//
	return &basicSongsService{songRepo: songRepo}
}

// New returns a SongsService with all of the expected middleware wired in.
func New(middleware []Middleware) SongsService {
	var svc SongsService = NewBasicSongsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicSongsService) GetAll(ctx context.Context) ([]model.Song, error) {

	return b.songRepo.FindAll()
}
