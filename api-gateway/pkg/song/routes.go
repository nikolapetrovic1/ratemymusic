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
	routes.GET("/:id", svc.FindOne)
	routes.GET("/musician/:id", svc.FindByMusician)
	routes.GET("/album/:id", svc.FindByAlbum)
	routes.GET("/search", svc.Search)
	routes.POST("/",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"MODERATOR"})
		}, svc.CreateSong)
	routes.PUT("/", func(c *gin.Context) {
		a.AuthRequired(c, []string{"MODERATOR"})
	}, svc.UpdateSong)
	routes.DELETE("/:id", func(c *gin.Context) {
		a.AuthRequired(c, []string{"MODERATOR"})
	}, svc.DeleteSong)
	routes.GET("/stream/:id", svc.Stream)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
func (svc *ServiceClient) FindByAlbum(ctx *gin.Context) {
	routes.FindByAlbum(ctx, svc.Client)
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
func (svc *ServiceClient) Search(ctx *gin.Context) {
	routes.Search(ctx, svc.Client)
}
func (svc *ServiceClient) Stream(ctx *gin.Context) {
	routes.Stream(ctx, svc.Client)
}
