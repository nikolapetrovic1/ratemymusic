package client

import (
	"context"
	"fmt"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"google.golang.org/grpc"
)

type MusicianServiceClient struct {
	Client pb.MusicianServiceClient
}

func InitMusicianServiceClient(url string) MusicianServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := MusicianServiceClient{
		Client: pb.NewMusicianServiceClient(cc),
	}

	return c
}

func (c *MusicianServiceClient) FindOne(musicianId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: musicianId,
	}

	return c.Client.FindOne(context.Background(), req)
}
