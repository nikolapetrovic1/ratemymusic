package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"net/http"
)

type MusicianUpdateBody struct {
	ID           int64  `json:"id"`
	MusicianName string `json:"musician_name"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}

func Update(ctx *gin.Context, c pb.MusicianServiceClient) {

	musicianData := MusicianUpdateBody{}

	if err := ctx.BindJSON(&musicianData); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateMusician(context.Background(), &pb.MusicianData{
		Id:           musicianData.ID,
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
