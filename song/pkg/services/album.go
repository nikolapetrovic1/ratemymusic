package services

import (
	"context"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/utils"
	"net/http"
)

//
//FindOne(context.Context, *IDRequest) (*BasicResponse, error)
//FindByMusician(context.Context, *FindByMusicianRequest) (*BasicResponse, error)
//CreateAlbum(context.Context, *AlbumData) (*BasicResponse, error)
//UpdateAlbum(context.Context, *AlbumData) (*BasicResponse, error)
//DeleteAlbum(context.Context, *IDRequest) (*DeleteResponse, error)
//SearchAlbum(context.Context, *SearchRequest) (*FindAllResponse, error)

type AlbumServer struct {
	pb.UnimplementedAlbumServiceServer
	Repo        db.Handler
	MusicianSvc client.MusicianServiceClient
}

func (s AlbumServer) FindOne(context context.Context, request *pb.IDRequest) (*pb.BasicResponse, error) {
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
