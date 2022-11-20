package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"net/http"
)

type FavoriteBody struct {
	Id     int64  `json:"id"`
	Kind   string `json:"kind"`
	KindId int64  `json:"kindId"`
	Name   string `json:"name"`
}

func Favorite(ctx *gin.Context, c pb.CollectionServiceClient) {

	request := FavoriteBody{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := ctx.GetInt64("userId")
	res, err := c.AddToFavorite(context.Background(), &pb.FavoriteData{
		Id:       request.Id,
		UserId:   userId,
		Kind:     request.Kind,
		KindId:   request.KindId,
		Name:     request.Name,
		Favorite: true,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
