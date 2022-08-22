package services

import (
	"context"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/utils"
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

func (s *Server) CreateSong(context context.Context, songRequest *pb.SongData) (*pb.CreateSongResponse, error) {
	var song models.Song

	song = utils.MapSongRequestSong(songRequest)
	s.Repo.DB.Create(&song)

	return &pb.CreateSongResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) UpdateSong(context context.Context, songRequest *pb.SongData) (*pb.CreateSongResponse, error) {
	s.Repo.DB.Save(utils.MapSongRequestSong(songRequest))

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

func (s *Server) SearchSong(context context.Context, request *pb.SearchRequest) (*pb.FindAllResponse, error) {
	return &pb.FindAllResponse{
		Status: http.StatusOK,
		Songs:  utils.MapListSongResponse(s.Repo.SearchSongs(request.Query)),
	}, nil
}

func (s *Server) FindByAlbum(context context.Context, request *pb.IdRequest) (*pb.FindByMusicianResponse, error) {
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
