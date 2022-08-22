package review

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/review"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ReviewServiceClient
}

func InitServiceClient(c *config.Config) pb.ReviewServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ReviewSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewReviewServiceClient(cc)
}
