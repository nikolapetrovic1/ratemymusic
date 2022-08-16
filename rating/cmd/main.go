package main

import (
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/services"
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

	pb.RegisterRatingServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
