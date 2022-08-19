package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"net/http"
	"strconv"
)

func FindByUserAlbum(ctx *gin.Context, c pb.RatingServiceClient) {
	album, err := strconv.ParseInt(ctx.Param("album"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.FindByUserAlbum(context.Background(), &pb.UserAlbumRequest{
		UserId:  ctx.GetInt64("userId"),
		AlbumId: album,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
