package song

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.SongServiceClient
}

func InitServiceClient(c *config.Config) pb.SongServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.SongSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewSongServiceClient(cc)
}
