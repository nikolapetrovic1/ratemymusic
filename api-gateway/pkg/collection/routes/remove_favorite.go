package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"net/http"
	"strconv"
)

func RemoveFavorite(ctx *gin.Context, c pb.CollectionServiceClient) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.GetInt64("userId")
	res, err := c.RemoveFavorite(context.Background(), &pb.DeleteRequest{
		Id:     id,
		UserId: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
