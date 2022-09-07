package routes

import (
	"context"
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
	value, err := stream.Recv()
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.Header("Content-Type", "audio/mpeg")
	//ctx.Header("Content-Disposition", `attachment;filename="music.mp3"`)

	//ctx.Header("Content-Type", "image/png")
	//ctx.Header("Content-Disposition", `attachment;filename="img.png"`)
	ctx.Data(http.StatusOK, "application/octet-stream", value.Data)
}
