package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"net/http"
)

type MusicianCreateBody struct {
	ID           int64  `json:"id"`
	MusicianName string `json:"musician_name"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}

func Create(ctx *gin.Context, c pb.MusicianServiceClient) {

	musicianData := MusicianCreateBody{}

	if err := ctx.BindJSON(&musicianData); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateMusician(context.Background(), &pb.MusicianData{
		Id:           0,
		MusicianName: musicianData.MusicianName,
		Name:         musicianData.Name,
		Surname:      musicianData.Surname,
		Songs:        nil,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
