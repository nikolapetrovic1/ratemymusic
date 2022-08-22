package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"net/http"
)

type AlbumRequestBody struct {
	Name       string `json:"name"`
	Duration   int32  `json:"duration"`
	MusicianID int64  `json:"musician_id"`
}

func Create(ctx *gin.Context, c pb.AlbumServiceClient) {

	request := AlbumRequestBody{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateAlbum(context.Background(), &pb.AlbumData{
		Id:         0,
		Name:       request.Name,
		MusicianID: request.MusicianID,
		Genres:     nil,
		Songs:      nil,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
