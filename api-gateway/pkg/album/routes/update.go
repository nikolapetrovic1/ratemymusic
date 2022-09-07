package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"net/http"
)

type AlbumUpdateRequestBody struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Duration   int32  `json:"duration"`
	MusicianID int64  `json:"musician_id"`
}

func Update(ctx *gin.Context, c pb.AlbumServiceClient) {

	req := AlbumUpdateRequestBody{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Update(context.Background(), &pb.AlbumData{
		Id:         req.ID,
		Name:       req.Name,
		MusicianID: req.MusicianID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
