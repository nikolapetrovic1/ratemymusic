package rating

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/rating/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/rating")

	routes.GET("user",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.FindByUser)
	routes.GET("user/song/:song",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.FindByUserSong)
	routes.GET("user/album/:album",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.FindByUserAlbum)
	routes.GET("song/:id", svc.FindBySong)
	routes.GET("album/:id", svc.FindByAlbum)
	routes.POST("album", svc.RateAlbum)
	routes.POST("song", func(c *gin.Context) {
		a.AuthRequired(c, []string{"USER"})
	}, svc.RateSong)
}

func (svc *ServiceClient) FindByUser(ctx *gin.Context) {
	routes.FindByUser(ctx, svc.Client)
}
func (svc *ServiceClient) FindBySong(ctx *gin.Context) {
	routes.FindBySong(ctx, svc.Client)
}
func (svc *ServiceClient) FindByAlbum(ctx *gin.Context) {
	routes.FindByAlbum(ctx, svc.Client)
}
func (svc *ServiceClient) FindByUserSong(ctx *gin.Context) {
	routes.FindByUserSong(ctx, svc.Client)
}
func (svc *ServiceClient) FindByUserAlbum(ctx *gin.Context) {
	routes.FindByUserAlbum(ctx, svc.Client)
}
func (svc *ServiceClient) RateAlbum(ctx *gin.Context) {
	routes.RateAlbum(ctx, svc.Client)
}
func (svc *ServiceClient) RateSong(ctx *gin.Context) {
	routes.RateSong(ctx, svc.Client)
}
