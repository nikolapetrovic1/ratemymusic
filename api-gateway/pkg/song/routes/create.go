package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"net/http"
)

type SongRequestBody struct {
	Name       string `json:"name"`
	Duration   int32  `json:"duration"`
	MusicianID int64  `json:"musician_id"`
}

func Create(ctx *gin.Context, c pb.SongServiceClient) {

	songRequest := SongRequestBody{}

	if err := ctx.BindJSON(&songRequest); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateSong(context.Background(), &pb.SongRequest{
		Name:       songRequest.Name,
		Duration:   songRequest.Duration,
		MusicianID: songRequest.MusicianID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
