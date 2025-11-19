package controllers

import (
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.IServiceRegistry
}

type IProductController interface {
	FindAllPrabayar(*gin.Context)
	FindAllPascabayar(*gin.Context)
	FindByProductCode(*gin.Context)
}

func NewBrandController(service services.IServiceRegistry) IProductController {
	return &ProductController{service: service}
}

func (b *ProductController) FindAllPrabayar(ctx *gin.Context) {
	user, ok := helper.GetAuthUser(ctx)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  ctx,
		})
		return
	}

	product, err := b.service.Product().FindAllPrabayar(ctx.Request.Context(), user.Role)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Err: err,
			Gin: ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Data: product,
		Gin:  ctx,
	})
}

func (b *ProductController) FindAllPascabayar(ctx *gin.Context) {
	user, ok := helper.GetAuthUser(ctx)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  ctx,
		})
		return
	}

	product, err := b.service.Product().FindAllPascabayar(ctx.Request.Context(), user.Role)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Err: err,
			Gin: ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Data: product,
		Gin:  ctx,
	})
}

func (b *ProductController) FindByProductCode(ctx *gin.Context) {
	user, ok := helper.GetAuthUser(ctx)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  ctx,
		})
		return
	}

	product, err := b.service.Product().FindByProductCode(ctx.Request.Context(), ctx.Param("product_code"), user.Role)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Err: err,
			Gin: ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Data: product,
		Gin:  ctx,
	})
}
