package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type BillingRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type IBillingRoute interface {
	Run()
}

func NewBillingRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) IBillingRoute {
	return &BillingRoute{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *BillingRoute) Run() {
	group := r.group.Group("/billing")
	group.POST("inquiry", r.middleware.Authenticate(), r.controller.BillingController().Inquiry)
	group.POST("pay", r.middleware.Authenticate(), r.controller.BillingController().Pay)
}
