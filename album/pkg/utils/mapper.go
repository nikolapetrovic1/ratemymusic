package utils

import (
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/models"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
)

func MapAlbumToAlbumResponse(album *models.Album) *pb.AlbumData {
	return &pb.AlbumData{
		Id:         album.ID,
		Name:       album.Name,
		MusicianID: album.MusicianID,
	}
}
func MapAlbumListToAlbumResponse(albums []models.Album) []*pb.AlbumData {
	var albumResponse []*pb.AlbumData
	for _, album := range albums {
		albumResponse = append(albumResponse, MapAlbumToAlbumResponse(&album))
	}
	return albumResponse
}

func MapAlbumDataToAlbum(albumData *pb.AlbumData) models.Album {
	return models.Album{
		ID:         albumData.Id,
		Name:       albumData.Name,
		MusicianID: albumData.MusicianID,
	}
}
