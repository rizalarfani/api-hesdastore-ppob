package controllers

import (
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
	"hesdastore/api-ppob/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service services.IServiceRegistry
}

type IAccountController interface {
	FindBalanceUser(*gin.Context)
}

func NewAccountController(service services.IServiceRegistry) IAccountController {
	return &AccountController{service: service}
}

func (c *AccountController) FindBalanceUser(ctx *gin.Context) {
	v, _ := ctx.Get("authUser")
	user, _ := v.(*model.ApiUser)

	balance, err := c.service.Account().FindBalanceUser(ctx, user.Username)
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
		Data: dto.AccountResponse{
			Name:    balance.Name,
			Balance: balance.Balance,
		},
		Gin: ctx,
	})
}
