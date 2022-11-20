package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"net/http"
)

func FindByUserId(ctx *gin.Context, c pb.CollectionServiceClient) {

	id := ctx.GetInt64("userId")
	res, err := c.GetByUserId(context.Background(), &pb.IdRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
