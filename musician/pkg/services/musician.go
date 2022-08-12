package services

import (
	"context"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"github.com/nikolapetrovic1/ratemymusic/musician/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/musician/pkg/models"

	"net/http"
)

type Server struct {
	pb.UnimplementedMusicianServiceServer
	Repo db.Handler
}

func (s *Server) FindOne(context context.Context, request *pb.FindOneRequest) (*pb.FindOneResponse, error) {

	var musician models.Musician

	if result := s.Repo.DB.First(&musician, request.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   mapMusicianMusicianData(musician),
	}, nil

}

func mapMusicianMusicianData(musician models.Musician) *pb.MusicianData {
	return &pb.MusicianData{
		Id:           musician.ID,
		MusicianName: musician.MusicianName,
		Name:         musician.Name,
		Surname:      musician.Surname,
		Songs:        nil,
	}
}
