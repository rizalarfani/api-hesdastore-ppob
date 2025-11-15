package config

import (
	"fmt"
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/constants"
	"hesdastore/api-ppob/controllers"
	"hesdastore/api-ppob/middlewares"
	"hesdastore/api-ppob/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes(c controllers.IControllerRegistry, m *middlewares.AuthMiddleware) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.ApiResponse{
			Status:  constants.Error,
			Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.ApiResponse{
			Status:  constants.Success,
			Message: "Welcome to API PPOB Hesda Store",
		})
	})

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")

		c.Next()
	})

	group := router.Group("/api/v1")
	route := routes.NewRouteRegistry(c, group, m)
	route.Serve()

	return router
}
