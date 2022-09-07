package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth/routes"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.GET("/all", svc.FindAll)
	routes.GET("/ban/:id", svc.Ban)
	routes.GET("/unban/:id", svc.Unban)
	routes.GET("/:id", svc.FindOne)
	routes.GET("/delete/:id", svc.Delete)
	routes.PUT("/", svc.Update)
	return svc
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}

func (svc *ServiceClient) Ban(ctx *gin.Context) {
	routes.BanUser(ctx, svc.Client)
}

func (svc *ServiceClient) Unban(ctx *gin.Context) {
	routes.UnBanUser(ctx, svc.Client)
}
func (svc *ServiceClient) FindAll(ctx *gin.Context) {
	routes.FindAll(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindById(ctx, svc.Client)
}
func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}
