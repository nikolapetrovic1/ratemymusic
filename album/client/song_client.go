package client

import (
	"context"
	"fmt"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	song "github.com/nikolapetrovic1/ratemymusic/song/pkg/models"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/utils"
	"google.golang.org/grpc"
)

type SongServiceClient struct {
	Client pb.SongServiceClient
}

func InitSongClient(url string) SongServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := SongServiceClient{
		Client: pb.NewSongServiceClient(cc),
	}

	return c
}

func (c *SongServiceClient) FindByAlbum(albumId int64) ([]song.Song, error) {
	req := &pb.IdRequest{
		Id: albumId,
	}
	response, err := c.Client.FindByAlbum(context.Background(), req)
	songs := utils.MapListSongDataSong(response.Songs)
	return songs, err
}
