package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type IRouterRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) IRouterRegistry {
	return &Registry{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *Registry) Serve() {
	r.brandRoute().Run()
	r.accountRoute().Run()
	r.productRoute().Run()
	r.transactionRoute().Run()
	r.billingRoute().Run()
	r.webhookRoute().Run()
}

func (r *Registry) brandRoute() IBrandRoute {
	return NewBrandRoute(r.controller, r.group, r.middleware)
}

func (r *Registry) accountRoute() IAccountRoute {
	return NewAccountRoute(r.controller, r.group, r.middleware)
}

func (r *Registry) productRoute() IProductRoute {
	return NewProductRoute(r.controller, r.group, r.middleware)
}

func (r *Registry) transactionRoute() ITransactionRoute {
	return NewTransactionRoute(r.controller, r.group, r.middleware)
}

func (r *Registry) billingRoute() IBillingRoute {
	return NewBillingRoute(r.controller, r.group, r.middleware)
}

func (r *Registry) webhookRoute() IWebhookRoute {
	return NewWebhookRoute(r.controller, r.group, r.middleware)
}
