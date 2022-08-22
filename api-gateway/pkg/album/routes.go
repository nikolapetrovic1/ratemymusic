package album

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/album/routes"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	//a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/album")

	//routes.GET("/:id", svc.FindOne)
	routes.GET("/musician/:id", svc.FindByMusician)
	routes.GET("/:id", svc.FindById)
	routes.GET("/search", svc.Search)
	routes.POST("/", svc.Create)
	//routes.PUT("/", svc.UpdateSong)
	//routes.DELETE("/:id", svc.DeleteSong)

}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}
func (svc *ServiceClient) FindByMusician(ctx *gin.Context) {
	routes.FindByMusician(ctx, svc.Client)
}
func (svc *ServiceClient) FindById(ctx *gin.Context) {
	routes.FindById(ctx, svc.Client)
}
func (svc *ServiceClient) Search(ctx *gin.Context) {
	routes.Search(ctx, svc.Client)
}
