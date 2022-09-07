package review

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/review/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/review")

	routes.POST("/",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.CreateReview)
	routes.GET("/song/:id", svc.FindBySong)
	routes.GET("/album/:id", svc.FindByAlbum)
	routes.GET("/user/:id", svc.FindByUserParam)
	routes.GET("/user",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUser)
	routes.GET("/user/song/:id",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserSong)
	routes.GET("/user/album/:id",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		}, svc.FindByUserAlbum)
}

func (svc *ServiceClient) FindByUserParam(ctx *gin.Context) {
	routes.FindByUserParam(ctx, svc.Client)
}
func (svc *ServiceClient) CreateReview(ctx *gin.Context) {
	routes.CreateReview(ctx, svc.Client)
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
