package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"net/http"
)

func FindByUserIdKind(ctx *gin.Context, c pb.CollectionServiceClient, kind string) {

	id := ctx.GetInt64("userId")
	res, err := c.GetByUserIdKind(context.Background(), &pb.KindRequest{
		Id: id, Kind: kind,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
