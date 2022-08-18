package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"net/http"
)

func Search(ctx *gin.Context, c pb.MusicianServiceClient) {
	query := ctx.DefaultQuery("query", "")

	res, err := c.SearchMusician(context.Background(), &pb.SearchRequest{
		Query: query,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
