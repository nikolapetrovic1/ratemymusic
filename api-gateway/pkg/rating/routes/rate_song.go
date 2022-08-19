package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"net/http"
)

type RateSongBody struct {
	Id          int64   `json:"id"`
	UserId      int64   `json:"user_id"`
	RatingValue float32 `json:"rating_value"`
	SongId      int64   `json:"song_id"`
}

func RateSong(ctx *gin.Context, c pb.RatingServiceClient) {
	rateSongBody := RateSongBody{}

	if err := ctx.BindJSON(&rateSongBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rateSongBody.UserId = ctx.GetInt64("userId")
	res, err := c.RateSong(context.Background(), &pb.RateRequest{
		Id:          rateSongBody.Id,
		RatingValue: rateSongBody.RatingValue,
		UserId:      rateSongBody.UserId,
		RatingType: &pb.RateRequest_SongId{
			SongId: rateSongBody.SongId,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
