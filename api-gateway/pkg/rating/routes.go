package rating

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/rating/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	_ = auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/rating")
	//routes.Use(a.AuthRequired)
	routes.GET("user/:id", svc.FindByUser)
	routes.GET("song/:id", svc.FindBySong)
	routes.GET("album/:id", svc.FindByAlbum)
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
