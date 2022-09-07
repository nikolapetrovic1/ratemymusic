package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"net/http"
)

type SongUpdateBody struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Duration   int32  `json:"duration"`
	MusicianID int64  `json:"musician_id"`
	AlbumId    int64  `json:"album_id"`
}

func Update(ctx *gin.Context, c pb.SongServiceClient) {

	songRequest := SongUpdateBody{}

	if err := ctx.BindJSON(&songRequest); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateSong(context.Background(), &pb.SongData{
		Id:         songRequest.Id,
		Name:       songRequest.Name,
		Duration:   songRequest.Duration,
		MusicianID: songRequest.MusicianID,
		AlbumId:    songRequest.AlbumId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
