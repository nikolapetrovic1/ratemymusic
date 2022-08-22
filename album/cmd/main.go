package main

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/album/client"
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/album/pkg/services"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/album"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	songSvc := client.InitSongClient(c.SongSvcUrl)
	fmt.Println("Song Svc on", c.Port)

	s := services.Server{
		Repo:    h,
		SongSvc: songSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAlbumServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
