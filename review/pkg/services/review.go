package services

import (
	"context"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/review"
	"github.com/nikolapetrovic1/ratemymusic/review/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/review/pkg/models"
	"net/http"
)

//FindByUser(context.Context, *IdRequest) (*ReviewResponse, error)
//FindBySong(context.Context, *IdRequest) (*ReviewResponse, error)
//FindByAlbum(context.Context, *IdRequest) (*ReviewResponse, error)
//FindByUserSong(context.Context, *UserSongRequest) (*ReviewData, error)
//FindByUserAlbum(context.Context, *UserAlbumRequest) (*ReviewData, error)
//CreateReview(context.Context, *ReviewData) (*ReviewData, error)

type Server struct {
	pb.UnimplementedReviewServiceServer
	Repo db.Handler
}

func (s Server) FindByUser(context context.Context, request *pb.IdRequest) (*pb.ReviewResponse, error) {
	var reviews []models.Review

	if result := s.Repo.DB.Where(&models.Review{UserId: request.Id}).Find(&reviews); result.Error != nil {
		return &pb.ReviewResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.ReviewResponse{
		Status:  http.StatusOK,
		Reviews: mapListReviewToReviewData(reviews),
	}, nil
}

func (s Server) FindBySong(context context.Context, request *pb.IdRequest) (*pb.ReviewResponse, error) {
	var reviews []models.Review

	if result := s.Repo.DB.Where(&models.Review{TypeId: request.Id, Type: "song"}).Find(&reviews); result.Error != nil {
		return &pb.ReviewResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.ReviewResponse{
		Status:  http.StatusOK,
		Reviews: mapListReviewToReviewData(reviews),
	}, nil
}
func (s Server) FindByAlbum(context context.Context, request *pb.IdRequest) (*pb.ReviewResponse, error) {
	var reviews []models.Review

	if result := s.Repo.DB.Where(&models.Review{TypeId: request.Id, Type: "album"}).Find(&reviews); result.Error != nil {
		return &pb.ReviewResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.ReviewResponse{
		Status:  http.StatusOK,
		Reviews: mapListReviewToReviewData(reviews),
	}, nil
}

func (s Server) FindByUserSong(context context.Context, request *pb.UserSongRequest) (*pb.ReviewData, error) {
	var review models.Review

	if result := s.Repo.DB.Where(&models.Review{UserId: request.UserId, TypeId: request.SongId, Type: "song"}).Order("updated_at").Find(&review); result.Error != nil {
		return nil, result.Error
	}
	return mapReviewToReviewData(review), nil
}

func (s Server) FindByUserAlbum(context context.Context, request *pb.UserAlbumRequest) (*pb.ReviewData, error) {
	var review models.Review

	if result := s.Repo.DB.Where(&models.Review{UserId: request.UserId, TypeId: request.AlbumId, Type: "album"}).Find(&review); result.Error != nil {
		return nil, result.Error
	}
	return mapReviewToReviewData(review), nil
}

func (s Server) CreateReview(context context.Context, request *pb.ReviewData) (*pb.ReviewData, error) {

	if s.checkIfReviewExists("song", request.UserId, request.TypeId) {
		s.Repo.DB.Save(mapReviewDataToReview(request))
		return request, nil
	}

	if s.checkIfReviewExists("album", request.UserId, request.TypeId) {
		s.Repo.DB.Save(mapReviewDataToReview(request))
		return request, nil
	}
	s.Repo.DB.Create(mapReviewDataToReview(request))

	return request, nil
}

func (s Server) checkIfReviewExists(Type string, userId int64, typeId int64) bool {
	if Type == "song" {
		review, _ := s.FindByUserSong(nil, &pb.UserSongRequest{
			UserId: userId,
			SongId: typeId,
		})
		if review == nil {
			return false
		}
	}
	if Type == "album" {
		review, _ := s.FindByUserAlbum(nil, &pb.UserAlbumRequest{
			UserId:  userId,
			AlbumId: typeId,
		})
		if review == nil {
			return false
		}
	}
	return true
}
