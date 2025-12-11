package controllers

import (
	error2 "hesdastore/api-ppob/common/error"
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/common/response"
	errConstant "hesdastore/api-ppob/constants/error"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BillingController struct {
	service services.IServiceRegistry
}

type IBillingController interface {
	Inquiry(*gin.Context)
	Pay(*gin.Context)
}

func NewBillingController(service services.IServiceRegistry) IBillingController {
	return &BillingController{service: service}
}

func (c *BillingController) Inquiry(g *gin.Context) {
	user, ok := helper.GetAuthUser(g)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  g,
		})
		return
	}

	var (
		request dto.InquiryBillRequest
		ctx     = g.Request.Context()
	)

	err := g.ShouldBindJSON(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  g,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(&request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := error2.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     errConstant.ErrValidatioin,
			Message: &errMessage,
			Errors:  errResponse,
			Gin:     g,
		})
		return
	}

	inquiry, err := c.service.Billing().Inquiry(ctx, &request, user)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  g,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: inquiry,
		Gin:  g,
	})
}

func (controller *BillingController) Pay(c *gin.Context) {
	user, ok := helper.GetAuthUser(c)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  c,
		})
		return
	}

	var (
		request dto.PayBillRequest
		ctx     = c.Request.Context()
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(&request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := error2.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Err:     errConstant.ErrValidatioin,
			Message: &errMessage,
			Errors:  errResponse,
			Gin:     c,
		})
		return
	}

	pay, err := controller.service.Billing().PayBill(ctx, &request, user)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: pay,
		Gin:  c,
	})
}
