package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"net/http"
)

func Search(ctx *gin.Context, c pb.AlbumServiceClient) {
	query := ctx.DefaultQuery("query", "")

	res, err := c.SearchAlbum(context.Background(), &pb.SearchRequest{
		Query: query,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
