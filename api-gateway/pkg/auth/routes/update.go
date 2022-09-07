package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/auth"
	"net/http"
)

type UpdateRequestBody struct {
	Email    string `json:"email"`
	NewEmail string `json:"new_email"`
}

func Update(ctx *gin.Context, c pb.AuthServiceClient) {
	b := UpdateRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Update(context.Background(), &pb.EmailUpdateRequest{
		Email:    b.Email,
		NewEmail: b.NewEmail,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
