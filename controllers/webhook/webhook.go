package controllers

import (
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	error2 "hesdastore/api-ppob/common/error"
	errConstant "hesdastore/api-ppob/constants/error"
)

type WebhookController struct {
	service services.IServiceRegistry
}

type IWebhookController interface {
	RetryWebhook(c *gin.Context)
}

func NewWebhookController(service services.IServiceRegistry) IWebhookController {
	return &WebhookController{
		service: service,
	}
}

func (controller *WebhookController) RetryWebhook(c *gin.Context) {
	var (
		request dto.RetryWebhookRequest
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

	err = controller.service.Webhook().RetryWebhook(ctx, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusInternalServerError,
			Err:  err,
			Gin:  c,
		})
		return
	}

	successMsg := "Webhook retry success"
	response.HttpResponse(response.ParamHTTPResp{
		Code:    http.StatusOK,
		Message: &successMsg,
		Gin:     c,
	})
}
