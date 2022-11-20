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

	cc, err := grpc.Dial(c.SongSvcUrl, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20)))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewSongServiceClient(cc)
}
