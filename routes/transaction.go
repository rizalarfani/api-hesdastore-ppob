package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type TransactionRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type ITransactionRoute interface {
	Run()
}

func NewTransactionRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) ITransactionRoute {
	return &TransactionRoute{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *TransactionRoute) Run() {
	group := r.group.Group("/transaction", r.middleware.Authenticate())
	group.POST("/order", r.controller.TransactionController().Order)
}
