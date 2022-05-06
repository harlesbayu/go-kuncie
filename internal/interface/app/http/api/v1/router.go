package v1

import (
	"github.com/harlesbayu/kuncie/internal/interface/app/http/api/v1/web"
	"github.com/harlesbayu/kuncie/internal/interface/container"
	"github.com/labstack/echo/v4"
)

type Router struct {
	V1Group        *echo.Group
	ProductHandler *web.ProductHandler
}

func NewRouter(server *echo.Echo, cont *container.Container) *Router {
	return &Router{
		V1Group:        server.Group("/v1"),
		ProductHandler: web.NewSubmissionHandler(cont.ProductService),
	}
}

func (r *Router) RegisterRoutes() {
	r.V1Group.GET("/products", r.ProductHandler.ListProductHandler)
	r.V1Group.POST("/scan-product", r.ProductHandler.ScanProductHandler)
}
