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
	//routes.Use(a.AuthRequired)
	routes.GET("/:id", a.AuthRequired, svc.FindOne)
	routes.GET("/musician/:id", svc.FindByMusician)
	routes.GET("/search", svc.Search)
	routes.POST("/", a.AuthRequired, svc.CreateSong)
	routes.PUT("/", a.AuthRequired, svc.UpdateSong)
	routes.DELETE("/:id", a.AuthRequired, svc.DeleteSong)
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
func (svc *ServiceClient) Search(ctx *gin.Context) {
	routes.Search(ctx, svc.Client)
}
