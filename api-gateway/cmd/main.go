package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/musician"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/rating"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/song"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)

	song.RegisterRoutes(r, &c, &authSvc)
	musician.RegisterRoutes(r, &c, &authSvc)
	rating.RegisterRoutes(r, &c, &authSvc)
	err = r.Run(c.Port)
	if err != nil {
		return
	}
}
