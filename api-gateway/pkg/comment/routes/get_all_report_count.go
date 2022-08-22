package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/comment"
	"net/http"
)

func GetAllByReportCount(ctx *gin.Context, c pb.CommentsClient) {

	res, err := c.GetAllByReportCount(context.Background(), &pb.IdRequest{
		Id: 0,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
