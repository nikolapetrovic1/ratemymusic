package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/review"
	"net/http"
)

type ReviewBody struct {
	ID     int64  `json:"id"`
	TypeId int64  `json:"type_id"`
	Type   string `json:"type"`
	Text   string `json:"text"`
}

func CreateReview(ctx *gin.Context, c pb.ReviewServiceClient) {

	request := ReviewBody{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateReview(context.Background(), &pb.ReviewData{
		Id:     request.ID,
		UserId: ctx.GetInt64("userId"),
		TypeId: request.TypeId,
		Text:   request.Text,
		Type:   request.Type,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
