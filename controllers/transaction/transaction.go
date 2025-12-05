package controllers

import (
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/constants"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/services"
	"io"
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
	GetHistory(*gin.Context)
	Order(*gin.Context)
	Webhooks(*gin.Context)
}

func NewBrandController(service services.IServiceRegistry) ITransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) GetHistory(ctx *gin.Context) {
	user, ok := helper.GetAuthUser(ctx)
	if !ok {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusUnauthorized,
			Gin:  ctx,
		})
		return
	}

	trxId := ctx.Params.ByName("transaction_id")

	historys, err := c.service.Transaction().GetHistory(ctx.Request.Context(), trxId, user.UserID)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusNotFound,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: historys,
		Gin:  ctx,
	})
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

func (c *TransactionController) Webhooks(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusMethodNotAllowed,
			Gin:  ctx,
		})
		return
	}

	bodyBytes, _ := io.ReadAll(ctx.Request.Body)

	xHubSignature := ctx.GetHeader(constants.XHubSignatureDigiflazz)
	if xHubSignature == "" {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Gin:  ctx,
		})
		return
	}

	err := c.service.Transaction().Webhooks(ctx, xHubSignature, bodyBytes)
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
		Gin:  ctx,
	})
}
