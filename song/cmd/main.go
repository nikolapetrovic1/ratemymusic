package main

import (
	"fmt"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/song"
	"github.com/nikolapetrovic1/ratemymusic/song/client"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/song/pkg/services"
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
	musicianSvc := client.InitMusicianServiceClient(c.MusicianSVCUrl)
	fmt.Println("Song Svc on", c.Port)

	s := services.Server{
		Repo:        h,
		MusicianSvc: musicianSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSongServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
