package utils

import (
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	pb1 "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
)

func MapAlbumToAlbumResponse(album *models.Album) *pb.AlbumData {
	return &pb.AlbumData{
		Id:         album.ID,
		Name:       album.Name,
		MusicianID: album.MusicianID,
	}
}

func MapListSongResponse(songs []models.Song) []*pb1.SongResponse {
	var songResponse []*pb1.SongResponse
	for _, song := range songs {
		songResponse = append(songResponse, MapSongResponse(song))
	}
	return songResponse
}

func MapSongResponse(song models.Song) *pb1.SongResponse {
	return &pb1.SongResponse{Id: song.ID, Name: song.Name, Duration: int32(song.Duration)}
}

func MapSongRequestSong(songRequest *pb1.SongRequest) *models.Song {
	return &models.Song{
		ID:         songRequest.Id,
		Name:       songRequest.Name,
		Duration:   int(songRequest.Duration),
		MusicianID: songRequest.MusicianID,
	}
}
