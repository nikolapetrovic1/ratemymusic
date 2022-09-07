package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/auth"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

func FindAll(ctx *gin.Context, c pb.AuthServiceClient) {

	res, err := c.GetAll(context.Background(), &emptypb.Empty{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
