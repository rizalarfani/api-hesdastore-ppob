package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type AccountRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type IAccountRoute interface {
	Run()
}

func NewAccountRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) IAccountRoute {
	return &AccountRoute{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *AccountRoute) Run() {
	group := r.group.Group("/account")
	group.GET("/saldo", r.middleware.Authenticate(), r.controller.AccountController().FindBalanceUser)
}
