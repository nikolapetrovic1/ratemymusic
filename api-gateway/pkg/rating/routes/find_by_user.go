package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"net/http"
)

func FindByUser(ctx *gin.Context, c pb.RatingServiceClient) {
	res, err := c.FindByUser(context.Background(), &pb.IdRequest{
		Id: ctx.GetInt64("userId"),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
