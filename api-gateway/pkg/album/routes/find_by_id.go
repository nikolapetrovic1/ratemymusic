package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"net/http"
	"strconv"
)

func FindById(ctx *gin.Context, c pb.AlbumServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.IDRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
