package song

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/song/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/song")
	routes.Use(a.AuthRequired)
	routes.GET("/:id", svc.FindOne)
	routes.GET("/musician/:id", svc.FindByMusician)
	routes.POST("/", svc.CreateSong)
	routes.PUT("/", svc.UpdateSong)
	routes.DELETE("/:id", svc.DeleteSong)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) FindByMusician(ctx *gin.Context) {
	routes.FindByMusician(ctx, svc.Client)
}

func (svc *ServiceClient) CreateSong(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateSong(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteSong(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}
