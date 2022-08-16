package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/comment"
	"net/http"
)

type CommentBody struct {
	ID       int64  `json:"id"`
	Comment  string `json:"comment"`
	UserId   int64  `json:"user_id"`
	ReviewId int64  `json:"review_id"`
}

func Create(ctx *gin.Context, c pb.CommentClient) {

	comment := CommentBody{}

	if err := ctx.BindJSON(&comment); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateComment(context.Background(), &pb.CommentRequest{
		Comment:  comment.Comment,
		UserId:   comment.UserId,
		ReviewId: comment.ReviewId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
