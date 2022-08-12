package services

import (
	"context"
	"fmt"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
	"net/http"
	"strconv"
)

type Server struct {
	pb.UnimplementedSongServiceServer
	Repo        db.Handler
	MusicianSvc client.MusicianServiceClient
}

func (s *Server) FindOne(ctx context.Context, request *pb.FindOneRequest) (*pb.FindOneResponse, error) {

	var song models.Song

	if result := s.Repo.DB.First(&song, request.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:   song.ID,
		Name: song.Name,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) FindByMusician(context context.Context, request *pb.FindByMusicianRequest) (*pb.FindByMusicianResponse, error) {
	_, err := s.MusicianSvc.FindOne(request.Id)
	if err != nil {
		return &pb.FindByMusicianResponse{
			Status: http.StatusNotFound,
			Error:  "Musician with ID:" + strconv.FormatInt(request.Id, 10) + " not found",
		}, nil
	}
	var songs []models.Song
	if result := s.Repo.DB.Where(&models.Song{MusicianID: request.Id}).Find(&songs); result.Error != nil {
		return &pb.FindByMusicianResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	fmt.Println(songs)
	return &pb.FindByMusicianResponse{
		Status: http.StatusOK,
		Songs:  mapListSongResponse(songs),
		Id:     request.Id,
	}, nil
}

func (s *Server) CreateSong(context context.Context, songRequest *pb.SongRequest) (*pb.CreateSongResponse, error) {
	s.Repo.DB.Create(mapSongRequestSong(songRequest))

	return &pb.CreateSongResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) UpdateSong(context context.Context, songRequest *pb.SongRequest) (*pb.CreateSongResponse, error) {
	s.Repo.DB.Save(mapSongRequestSong(songRequest))

	return &pb.CreateSongResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) DeleteSong(context context.Context, deleteRequest *pb.DeleteSongRequest) (*pb.CreateSongResponse, error) {
	s.Repo.DB.Delete(&models.Song{}, deleteRequest.Id)

	return &pb.CreateSongResponse{
		Status: http.StatusAccepted,
	}, nil
}

func mapListSongResponse(songs []models.Song) []*pb.SongResponse {
	var songResponse []*pb.SongResponse
	for _, song := range songs {
		songResponse = append(songResponse, mapSongResponse(song))
	}
	return songResponse
}
func mapSongResponse(song models.Song) *pb.SongResponse {
	return &pb.SongResponse{Id: song.ID, Name: song.Name, Duration: int32(song.Duration)}
}

func mapSongRequestSong(songRequest *pb.SongRequest) *models.Song {
	return &models.Song{
		ID:         songRequest.Id,
		Name:       songRequest.Name,
		Duration:   int(songRequest.Duration),
		MusicianID: songRequest.MusicianID,
	}
}
