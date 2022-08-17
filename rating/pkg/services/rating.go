package services

import (
	"context"
	"fmt"
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
	if result := s.Repo.DB.Where(&models.SongRating{UserID: 1}, request.Id).Find(&songRatings); result.Error != nil {
		return &pb.RatingResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	if result := s.Repo.DB.Where(&models.AlbumRating{UserID: 1}, request.Id).Find(&albumRatings); result.Error != nil {
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

	if result := s.Repo.DB.Where(&models.SongRating{SongID: 1}, request.Id).Find(&songRatings); result.Error != nil {
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

	if result := s.Repo.DB.Where(&models.AlbumRating{AlbumId: 1}, request.Id).Find(&albumRatings); result.Error != nil {
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
	_, err := s.FindReviewByUserSong(request.UserId, request.GetSongId())
	if err == nil {
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
	_, err := s.FindReviewByUserAlbum(request.UserId, request.GetAlbumId())
	fmt.Println(err)
	if err == nil {
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

func (s Server) FindReviewByUserSong(userId int64, songId int64) (models.SongRating, error) {
	var review models.SongRating
	if result := s.Repo.DB.Where(&models.SongRating{UserID: userId, SongID: songId}).Find(&review); result.Error != nil {
		return review, result.Error
	}
	return review, nil
}

func (s Server) FindReviewByUserAlbum(userId int64, albumId int64) (models.AlbumRating, error) {
	var review models.AlbumRating
	if result := s.Repo.DB.Where(&models.AlbumRating{UserID: userId, AlbumId: albumId}).Find(&review); result.Error != nil {
		return review, result.Error
	}
	return review, nil
}
