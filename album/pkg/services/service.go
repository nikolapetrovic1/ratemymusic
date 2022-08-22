package services

import (
	"context"
	"github.com/nikolapetrovic1/ratemymusic/album/client"
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/utils"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"net/http"
)

type Server struct {
	pb.UnimplementedAlbumServiceServer
	Repo    db.Handler
	SongSvc client.SongServiceClient
}

func (s Server) FindOne(context context.Context, request *pb.IDRequest) (*pb.BasicResponse, error) {
	var album models.Album

	if result := s.Repo.DB.First(&album, request.Id); result.Error != nil {
		return &pb.BasicResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.BasicResponse{
		Status: http.StatusOK,
		Data:   utils.MapAlbumToAlbumResponse(&album),
	}, nil
}

func (s *Server) SearchAlbum(context context.Context, request *pb.SearchRequest) (*pb.FindAllResponse, error) {

	var albums []*pb.AlbumData
	albums = utils.MapAlbumListToAlbumResponse(s.Repo.SearchAlbums(request.Query))
	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Albums: albums,
	}, nil
}

func (s *Server) FindByMusician(_ context.Context, request *pb.IDRequest) (*pb.FindAllResponse, error) {
	var albums []models.Album
	s.Repo.DB.Where(&models.Album{MusicianID: request.Id}).Find(&albums)
	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Albums: utils.MapAlbumListToAlbumResponse(albums),
	}, nil
}

func (s *Server) CreateAlbum(context context.Context, request *pb.AlbumData) (*pb.BasicResponse, error) {
	var album models.Album
	album = utils.MapAlbumDataToAlbum(request)

	s.Repo.DB.Create(&album)
	return &pb.BasicResponse{
		Status: http.StatusCreated,
	}, nil
}
