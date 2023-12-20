package router

import (
	"github.com/fadilahonespot/simple-api/server/handler"
	"github.com/fadilahonespot/simple-api/server/middleware"
	"github.com/labstack/echo/v4"
)

type DefaultRouter struct {
	ProductHandler *handler.ProductHandler
}

func (d *DefaultRouter) Validate() {
	if d.ProductHandler == nil {
		panic("product handler is nil")
	}
}

func (d *DefaultRouter) NewRouter(e *echo.Echo) *DefaultRouter {
	middleware.SetupMiddleware(e)
	
	e.POST("/products", d.ProductHandler.AddProduct)
	e.GET("/products", d.ProductHandler.GetListProduct)
	e.GET("/products/:productId", d.ProductHandler.GetProductDetail)
	e.PUT("/products/:productId", d.ProductHandler.UpdateProduct)
	e.DELETE("/products/:productId", d.ProductHandler.DeleteProduct)
	
	return d
}