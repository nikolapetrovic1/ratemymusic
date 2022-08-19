package services

import (
	"context"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"net/http"
)

type Server struct {
	pb.UnimplementedRatingServiceServer
	Repo        db.Handler
	MusicianSvc client.MusicianServiceClient
}

func (s Server) FindByUser(context context.Context, request *pb.IdRequest) (*pb.RatingResponse, error) {
	var songRatings []models.SongRating
	var albumRatings []models.AlbumRating
	if result := s.Repo.DB.Where(&models.SongRating{UserID: request.Id}).Find(&songRatings); result.Error != nil {
		return &pb.RatingResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	if result := s.Repo.DB.Where(&models.AlbumRating{UserID: request.Id}).Find(&albumRatings); result.Error != nil {
		return &pb.RatingResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.RatingResponse{
		Status:     http.StatusOK,
		RatingData: mapListRatingRatingData(songRatings, albumRatings),
	}, nil
}

func (s Server) FindBySong(context context.Context, request *pb.IdRequest) (*pb.RatingResponse, error) {
	var songRatings []models.SongRating

	if result := s.Repo.DB.Where(&models.SongRating{SongID: request.Id}).Find(&songRatings); result.Error != nil {
		return &pb.RatingResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.RatingResponse{
		Status:     http.StatusOK,
		RatingData: mapListSongRatingRatingData(songRatings),
	}, nil
}

func (s Server) FindByAlbum(context context.Context, request *pb.IdRequest) (*pb.RatingResponse, error) {
	var albumRatings []models.AlbumRating

	if result := s.Repo.DB.Where(&models.AlbumRating{AlbumId: request.Id}).Find(&albumRatings); result.Error != nil {
		return &pb.RatingResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.RatingResponse{
		Status:     http.StatusOK,
		RatingData: mapListAlbumRatingRatingData(albumRatings),
	}, nil
}
func (s Server) RateSong(context context.Context, request *pb.RateRequest) (*pb.RateResponse, error) {

	rating, err := s.FindByUserSong(context, &pb.UserSongRequest{
		UserId: request.UserId,
		SongId: request.GetSongId(),
	})
	if err == nil {
		request.Id = rating.Id
		s.Repo.DB.Save(mapRatingDataRatingSong(request))
		return &pb.RateResponse{
			Status: http.StatusOK,
		}, nil
	}
	s.Repo.DB.Create(mapRatingDataRatingSong(request))

	return &pb.RateResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Server) RateAlbum(context context.Context, request *pb.RateRequest) (*pb.RateResponse, error) {
	rating, err := s.FindByUserAlbum(context, &pb.UserAlbumRequest{
		UserId:  request.UserId,
		AlbumId: request.GetAlbumId(),
	})
	if err == nil {
		request.Id = rating.Id
		s.Repo.DB.Save(mapRatingDataRatingAlbum(request))
		return &pb.RateResponse{
			Status: http.StatusOK,
		}, nil
	}
	s.Repo.DB.Create(mapRatingDataRatingAlbum(request))

	return &pb.RateResponse{
		Status: http.StatusOK,
	}, nil
}

func (s Server) FindByUserAlbum(context context.Context, request *pb.UserAlbumRequest) (*pb.RatingData, error) {
	var rating models.AlbumRating
	if result := s.Repo.DB.Where(&models.AlbumRating{UserID: request.UserId, AlbumId: request.AlbumId}).Find(&rating); result.Error != nil {
		return nil, result.Error
	}
	return mapAlbumRatingRatingData(&rating), nil
}
func (s Server) FindByUserSong(context context.Context, request *pb.UserSongRequest) (*pb.RatingData, error) {
	var rating models.SongRating
	if result := s.Repo.DB.Where(&models.SongRating{UserID: request.UserId, SongID: request.SongId}).Find(&rating); result.Error != nil {
		return nil, result.Error
	}
	return mapSongRatingRatingData(&rating), nil
}
