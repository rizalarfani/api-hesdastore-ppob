package controllers

import (
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrandController struct {
	service services.IServiceRegistry
}

type IBrandController interface {
	FindAll(*gin.Context)
}

func NewBrandController(service services.IServiceRegistry) IBrandController {
	return &BrandController{service: service}
}

func (b *BrandController) FindAll(ctx *gin.Context) {
	brand, err := b.service.Brand().FindAll(ctx.Request.Context())
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: brand,
		Gin:  ctx,
	})
}
