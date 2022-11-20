package musician

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.MusicianServiceClient
}

func InitServiceClient(c *config.Config) pb.MusicianServiceClient {
	cc, err := grpc.Dial(c.MusicianSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewMusicianServiceClient(cc)
}
