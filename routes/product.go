package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type ProductRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type IProductRoute interface {
	Run()
}

func NewProductRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) IAccountRoute {
	return &ProductRoute{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *ProductRoute) Run() {
	group := r.group.Group("/product", r.middleware.Authenticate())
	group.GET("/prabayar", r.controller.ProductController().FindAllPrabayar)
	group.GET("/pascabayar", r.controller.ProductController().FindAllPascabayar)
	group.GET("/:product_code", r.controller.ProductController().FindByProductCode)
}
