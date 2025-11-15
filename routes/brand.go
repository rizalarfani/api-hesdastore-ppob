package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type BrandRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type IBrandRoute interface {
	Run()
}

func NewBrandRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) IBrandRoute {
	return &BrandRoute{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *BrandRoute) Run() {
	group := r.group.Group("/brand")
	group.GET("", r.middleware.Authenticate(), r.controller.BrandController().FindAll)
}
