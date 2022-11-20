package collection

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/collection/routes"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/collection")
	routes.GET("/user",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserId)
	routes.GET("/user/album",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserIdAlbum)
	routes.GET("/user/musician",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserIdAlbum)
	routes.GET("/user/song/",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserIdSong)
	routes.GET("/user/song/:kindId",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserIdSongId)
	routes.GET("/user/album/:kindId",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserIdAlbumId)
	routes.GET("/user/musician/:kindId",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserIdMusicianId)
	routes.POST("/favorite",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.Favorite)
	routes.GET("/unfavorite/:id",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.RemoveFavorite)
}

func (svc *ServiceClient) FindByUserId(ctx *gin.Context) {
	routes.FindByUserId(ctx, svc.Client)
}
func (svc *ServiceClient) FindByUserIdSong(ctx *gin.Context) {
	routes.FindByUserIdKind(ctx, svc.Client, "song")
}
func (svc *ServiceClient) FindByUserIdMusician(ctx *gin.Context) {
	routes.FindByUserIdKind(ctx, svc.Client, "musician")
}
func (svc *ServiceClient) FindByUserIdAlbum(ctx *gin.Context) {
	routes.FindByUserIdKind(ctx, svc.Client, "album")
}

func (svc *ServiceClient) FindByUserIdSongId(ctx *gin.Context) {
	routes.FindByUserIdKindId(ctx, svc.Client, "song")
}
func (svc *ServiceClient) FindByUserIdMusicianId(ctx *gin.Context) {
	routes.FindByUserIdKindId(ctx, svc.Client, "musician")
}
func (svc *ServiceClient) FindByUserIdAlbumId(ctx *gin.Context) {
	routes.FindByUserIdKindId(ctx, svc.Client, "album")
}

func (svc *ServiceClient) Favorite(ctx *gin.Context) {
	routes.Favorite(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveFavorite(ctx *gin.Context) {
	routes.RemoveFavorite(ctx, svc.Client)
}
