package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"net/http"
	"strconv"
)

func FindByUserSong(ctx *gin.Context, c pb.RatingServiceClient) {
	song, err := strconv.ParseInt(ctx.Param("song"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.FindByUserSong(context.Background(), &pb.UserSongRequest{
		UserId: ctx.GetInt64("userId"),
		SongId: song,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
