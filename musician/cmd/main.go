package main

import (
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/musician"
	"github.com/nikolapetrovic1/ratemymusic/musician/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/musician/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/musician/pkg/services"
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

	s := services.Server{
		Repo: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMusicianServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
