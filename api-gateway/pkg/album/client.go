package album

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AlbumServiceClient
}

func InitServiceClient(c *config.Config) pb.AlbumServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AlbumSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAlbumServiceClient(cc)
}
