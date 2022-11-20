package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
	"net/http"
	"strconv"
)

func FindByUserIdKindId(ctx *gin.Context, c pb.CollectionServiceClient, kind string) {
	kindId, err := strconv.ParseInt(ctx.Param("kindId"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	id := ctx.GetInt64("userId")
	res, err := c.GetByUserIdKindId(context.Background(), &pb.KindIdRequest{
		Id: id, Kind: kind, KindId: kindId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
