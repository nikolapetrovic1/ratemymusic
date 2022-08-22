package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/comment"
	"net/http"
	"strconv"
)

func Delete(ctx *gin.Context, c pb.CommentsClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.DeleteComment(context.Background(), &pb.IdRequest{
		Id: int32(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
