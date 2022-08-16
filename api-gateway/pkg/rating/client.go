package rating

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.RatingServiceClient
}

func InitServiceClient(c *config.Config) pb.RatingServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.RatingSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewRatingServiceClient(cc)
}
