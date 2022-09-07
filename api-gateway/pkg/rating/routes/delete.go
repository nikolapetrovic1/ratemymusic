package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"net/http"
	"strconv"
)

func Delete(ctx *gin.Context, c pb.RatingServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	Type := ctx.Param("type")
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id:   id,
		Type: Type,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
