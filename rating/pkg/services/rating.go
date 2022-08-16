package services

import (
	"context"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"net/http"
)

//FindByUser(context.Context, *IdRequest) (*RatingResponse, error)
//FindBySong(context.Context, *IdRequest) (*RatingResponse, error)
//FindByAlbum(context.Context, *IdRequest) (*RatingResponse, error)
//RateSong(context.Context, *RateRequest) (*RateResponse, error)
//RateAlbum(context.Context, *RateRequest) (*RateResponse, error)

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

func mapListRatingRatingData(ratingsSong []models.SongRating, ratingsAlbum []models.AlbumRating) []*pb.RatingData {
	var ratingDatas []*pb.RatingData
	for _, rating := range ratingsSong {
		ratingDatas = append(ratingDatas, mapSongRatingRatingData(&rating))
	}
	for _, rating := range ratingsAlbum {
		ratingDatas = append(ratingDatas, mapAlbumRatingRatingData(&rating))
	}
	return ratingDatas
}
func mapListSongRatingRatingData(ratingsSong []models.SongRating) []*pb.RatingData {
	var ratingDatas []*pb.RatingData

	for _, rating := range ratingsSong {
		ratingDatas = append(ratingDatas, mapSongRatingRatingData(&rating))
	}
	return ratingDatas
}

func mapListAlbumRatingRatingData(ratingsAlbum []models.AlbumRating) []*pb.RatingData {
	var ratingDatas []*pb.RatingData

	for _, rating := range ratingsAlbum {
		ratingDatas = append(ratingDatas, mapAlbumRatingRatingData(&rating))
	}
	return ratingDatas
}
func mapSongRatingRatingData(rating *models.SongRating) *pb.RatingData {
	return &pb.RatingData{
		RatingValue: float32(rating.RatingValue),
		UserId:      rating.UserID,
		RatingType: &pb.RatingData_SongId{
			SongId: rating.SongID,
		},
	}
}
func mapAlbumRatingRatingData(rating *models.AlbumRating) *pb.RatingData {
	return &pb.RatingData{
		RatingValue: float32(rating.RatingValue),
		UserId:      rating.UserID,
		RatingType: &pb.RatingData_AlbumId{
			AlbumId: rating.AlbumId,
		},
	}
}
