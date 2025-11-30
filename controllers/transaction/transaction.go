package controllers

import (
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/services"
	"net/http"

	error2 "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionController struct {
	service services.IServiceRegistry
}

type ITransactionController interface {
	Order(*gin.Context)
}

func NewBrandController(service services.IServiceRegistry) ITransactionController {
	return &TransactionController{service: service}
}

func (transaction *TransactionController) Order(c *gin.Context) {
	user, ok := helper.GetAuthUser(c)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  c,
		})
		return
	}

	var (
		request dto.TransactionOrderRequest
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

	order, err := transaction.service.Transaction().Order(ctx, &request, user)
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
		Data: order,
		Gin:  c,
	})
}
