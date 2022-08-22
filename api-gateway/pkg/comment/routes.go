package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/auth"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/comment/routes"
	"github.com/nikolapetrovic1/ratemymusic/api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/comment")
	routes.POST("/",
		func(c *gin.Context) {
			a.AuthRequired(c, []string{"USER"})
		},
		svc.Create)
	routes.GET("/review/:id", svc.FindByReview)
	routes.GET("/delete/:id", svc.Delete)
	routes.GET("/report/all", svc.GetAllByReportCount)
	routes.GET("/report/:id", svc.Report)

}

func (svc *ServiceClient) GetAllByReportCount(ctx *gin.Context) {
	routes.GetAllByReportCount(ctx, svc.Client)
}
func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}
func (svc *ServiceClient) FindByReview(ctx *gin.Context) {
	routes.FindByReview(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}
func (svc *ServiceClient) Report(ctx *gin.Context) {
	routes.Report(ctx, svc.Client)
}

//func (svc *ServiceClient) Update(ctx *gin.Context) {
//	routes.Update(ctx, svc.Client)
//}
//func (svc *ServiceClient) Delete(ctx *gin.Context) {
//	routes.Delete(ctx, svc.Client)
//}
//
//func (svc *ServiceClient) Search(ctx *gin.Context) {
//	routes.Delete(ctx, svc.Client)
//}
