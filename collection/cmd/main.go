package main

import (
	"fmt"
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/db"
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/services"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
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

	pb.RegisterCollectionServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	fmt.Println("Collection Svc on", c.Port)
}
