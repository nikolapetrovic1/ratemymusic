package services

import (
	"context"
	"errors"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/utils"
	"net/http"
	"strconv"
)

type Server struct {
	Repo        db.Handler
	MusicianSvc client.MusicianServiceClient
	pb.UnimplementedSongServiceServer
}

func (s *Server) StreamSong(_ context.Context, request *pb.FindOneRequest) (*pb.AudioSample, error) {
	var song models.Song
	if result := s.Repo.DB.First(&song, request.Id); result.Error != nil {
		return &pb.AudioSample{
			Data: nil,
		}, errors.New("song does not exist")
	}

	return &pb.AudioSample{
		Data: utils.LoadAudio(song.Audio),
	}, nil
}

func (s *Server) FindOne(_ context.Context, request *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var song models.Song

	if result := s.Repo.DB.First(&song, request.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   utils.MapSongResponse(song),
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
	return &pb.FindByMusicianResponse{
		Status: http.StatusOK,
		Songs:  utils.MapListSongResponse(songs),
		Id:     request.Id,
	}, nil
}

func (s *Server) CreateSong(_ context.Context, songRequest *pb.SongData) (*pb.CreateSongResponse, error) {

	var song models.Song

	song = utils.MapSongRequestSong(songRequest)
	s.Repo.DB.Create(&song)

	return &pb.CreateSongResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) UpdateSong(_ context.Context, songRequest *pb.SongData) (*pb.CreateSongResponse, error) {
	var song models.Song
	song = utils.MapSongRequestSong(songRequest)
	s.Repo.DB.Save(&song)

	return &pb.CreateSongResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) DeleteSong(_ context.Context, deleteRequest *pb.DeleteSongRequest) (*pb.CreateSongResponse, error) {
	s.Repo.DB.Delete(&models.Song{}, deleteRequest.Id)

	return &pb.CreateSongResponse{
		Status: http.StatusAccepted,
	}, nil
}

func (s *Server) SearchSong(_ context.Context, request *pb.SearchRequest) (*pb.FindAllResponse, error) {
	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Songs:  utils.MapListSongResponse(s.Repo.SearchSongs(request.Query, "Hip-Hop")),
	}, nil
}

func (s *Server) FindByAlbum(_ context.Context, request *pb.IdRequest) (*pb.FindByMusicianResponse, error) {
	var songs []models.Song
	if result := s.Repo.DB.Where(&models.Song{AlbumID: request.Id}).Find(&songs); result.Error != nil {
		return &pb.FindByMusicianResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.FindByMusicianResponse{
		Status: http.StatusOK,
		Songs:  utils.MapListSongResponse(songs),
		Id:     request.Id,
	}, nil
}

func (s *Server) MusicianExists(id int64) error {
	_, err := s.MusicianSvc.FindOne(id)
	return err
}
