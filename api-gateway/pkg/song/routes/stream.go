package routes

import (
	"context"
	b64 "encoding/base64"
	"github.com/gin-gonic/gin"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"net/http"
	"strconv"
)

func Stream(ctx *gin.Context, c pb.SongServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	stream, err := c.StreamSong(context.Background(), &pb.FindOneRequest{
		Id: id,
	})
	if err != nil {
		println(err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	encodedStream := b64.StdEncoding.EncodeToString(stream.Data)
	ctx.JSON(http.StatusOK, &encodedStream)
}
