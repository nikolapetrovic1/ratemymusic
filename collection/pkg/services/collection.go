package services

import (
	"context"
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/utils"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"net/http"
)

type Server struct {
	pb.UnimplementedCollectionServiceServer
	Repo db.Handler
}

func (s *Server) GetByUserIdKind(_ context.Context, request *pb.KindRequest) (*pb.BasicResponse, error) {

	var collection []models.Favorite

	if result := s.Repo.DB.Where(&models.Favorite{UserId: request.Id, Kind: request.Kind}).Find(&collection).Order("kind"); result.Error != nil {
		return &pb.BasicResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil

	}
	return &pb.BasicResponse{
		Status:    request.Id,
		Error:     "",
		Favorites: utils.MapListCollectionToCollectionData(collection),
	}, nil

}
func (s *Server) GetByUserId(_ context.Context, request *pb.IdRequest) (*pb.BasicResponse, error) {

	var collection []models.Favorite

	if result := s.Repo.DB.Where(&models.Favorite{UserId: request.Id}).Find(&collection).Order("kind"); result.Error != nil {
		return &pb.BasicResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil

	}
	return &pb.BasicResponse{
		Status:    http.StatusOK,
		Error:     "",
		Favorites: utils.MapListCollectionToCollectionData(collection),
	}, nil
}

func (s *Server) AddToFavorite(_ context.Context, request *pb.FavoriteData) (*pb.AddToFavoriteResponse, error) {

	var favorite models.Favorite

	//if result := s.Repo.DB.Where(&models.Favorite{UserId: request.Id}) result.Error != nil

	favorite = utils.MapFavoriteDataToFavorite(request)
	s.Repo.DB.Save(&favorite)

	return &pb.AddToFavoriteResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (s *Server) GetByUserIdKindId(_ context.Context, request *pb.KindIdRequest) (*pb.KindIdResponse, error) {
	var favorite models.Favorite

	if result := s.Repo.DB.Where(&models.Favorite{UserId: request.Id, Kind: request.Kind, KindId: request.KindId}).First(&favorite); result.Error != nil {
		return &pb.KindIdResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.KindIdResponse{
		Status:   http.StatusOK,
		Favorite: utils.MapFavoriteToFavoriteData(favorite),
	}, nil
}

func (s *Server) RemoveFavorite(_ context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	if result := s.Repo.DB.Delete(&models.Favorite{ID: request.Id, UserId: request.UserId}); result.Error != nil {
		return &pb.DeleteResponse{}, result.Error
	}
	return &pb.DeleteResponse{Info: "Removed from favorite "}, nil
}
