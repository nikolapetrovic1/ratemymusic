package utils

import (
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
)

func MapListSongResponse(songs []models.Song) []*pb.SongData {
	var songResponse []*pb.SongData
	for _, song := range songs {
		songResponse = append(songResponse, MapSongResponse(song))
	}
	return songResponse
}

func MapListSongDataSong(songsData []*pb.SongData) []models.Song {
	var songs []models.Song
	for _, song := range songsData {
		songs = append(songs, MapSongRequestSong(song))
	}
	return songs
}

func MapSongResponse(song models.Song) *pb.SongData {
	return &pb.SongData{
		Id:         song.ID,
		Name:       song.Name,
		Duration:   int32(song.Duration),
		MusicianID: song.MusicianID,
		Genres:     nil,
	}
}

func MapSongRequestSong(songRequest *pb.SongData) models.Song {
	return models.Song{
		ID:          songRequest.Id,
		Name:        songRequest.Name,
		Duration:    int(songRequest.Duration),
		MusicianID:  songRequest.MusicianID,
		Description: "",
		AlbumID:     songRequest.AlbumId,
	}
}
