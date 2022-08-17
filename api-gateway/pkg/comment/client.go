package comment

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/comment"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CommentsClient
}

func InitServiceClient(c *config.Config) pb.CommentsClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCommentsClient(cc)
}
