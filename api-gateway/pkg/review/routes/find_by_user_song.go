package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/review"
	"net/http"
	"strconv"
)

func FindByUserSong(ctx *gin.Context, c pb.ReviewServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindByUserSong(context.Background(), &pb.UserSongRequest{
		UserId: ctx.GetInt64("userId"),
		SongId: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}