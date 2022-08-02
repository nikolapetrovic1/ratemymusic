package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	model "github.com/nikolapetrovic1/ratemymusic/songs/pkg/model"
)

// Middleware describes a service middleware.
type Middleware func(SongsService) SongsService

type loggingMiddleware struct {
	logger log.Logger
	next   SongsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SongsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SongsService) SongsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (m0 model.Song, e1 error) {
	defer func() {
		l.logger.Log("method", "Get", "m0", m0, "e1", e1)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) GetAll(ctx context.Context) (m0 []model.Song, e1 error) {
	defer func() {
		l.logger.Log("method", "GetAll", "m0", m0, "e1", e1)
	}()
	return l.next.GetAll(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, song model.Song) (m0 model.Song, e1 error) {
	defer func() {
		l.logger.Log("method", "Add", "song", song, "m0", m0, "e1", e1)
	}()
	return l.next.Add(ctx, song)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "e0", e0)
	}()
	return l.next.Delete(ctx, id)
}
