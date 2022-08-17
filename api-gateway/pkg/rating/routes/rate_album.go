package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"net/http"
)

type RateAlbumBody struct {
	Id          int64   `json:"id"`
	UserId      int64   `json:"user_id"`
	RatingValue float32 `json:"rating_value"`
	AlbumId     int64   `json:"album_id"`
}

func RateAlbum(ctx *gin.Context, c pb.RatingServiceClient) {
	rateAlbumBody := RateAlbumBody{}

	if err := ctx.BindJSON(&rateAlbumBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.RateAlbum(context.Background(), &pb.RateRequest{
		Id:          rateAlbumBody.Id,
		RatingValue: rateAlbumBody.RatingValue,
		UserId:      rateAlbumBody.UserId,
		RatingType: &pb.RateRequest_AlbumId{
			AlbumId: rateAlbumBody.AlbumId,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
