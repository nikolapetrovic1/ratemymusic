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
	ReviewId int32  `json:"review_id"`
}

func Create(ctx *gin.Context, c pb.CommentsClient) {

	comment := CommentBody{}

	if err := ctx.BindJSON(&comment); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.CreateComment(context.Background(), &pb.CommentRequest{
		Comment:  comment.Comment,
		UserId:   int32(ctx.GetInt64("userId")),
		ReviewId: comment.ReviewId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
