package collection

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CollectionServiceClient
}

func InitServiceClient(c *config.Config) pb.CollectionServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CollectionSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCollectionServiceClient(cc)
}
