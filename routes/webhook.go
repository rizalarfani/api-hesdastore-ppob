package routes

import (
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"

	"github.com/gin-gonic/gin"
)

type WebhookRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	middleware *middlewares.AuthMiddleware
}

type IWebhookRoute interface {
	Run()
}

func NewWebhookRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, middleware *middlewares.AuthMiddleware) IWebhookRoute {
	return &WebhookRoute{
		controller: controller,
		group:      group,
		middleware: middleware,
	}
}

func (r *WebhookRoute) Run() {
	group := r.group.Group("/webhook")
	group.POST("/retry", r.middleware.Authenticate(), r.controller.WebhookController().RetryWebhook)
}
