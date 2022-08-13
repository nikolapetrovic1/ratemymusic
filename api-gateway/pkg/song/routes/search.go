package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"net/http"
)

func Search(ctx *gin.Context, c pb.SongServiceClient) {
	query := ctx.DefaultQuery("query", "")

	res, err := c.SearchSong(context.Background(), &pb.SearchRequest{
		Query: query,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
