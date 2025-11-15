package response

import (
	"hesdastore/api-ppob/constants"
	"net/http"

	"github.com/gin-gonic/gin"

	errConstant "hesdastore/api-ppob/constants/error"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
}

func HttpResponse(param ParamHTTPResp) {
	if param.Err == nil {
		param.Gin.JSON(param.Code, ApiResponse{
			Status:  constants.Success,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
		})
		return
	}

	message := errConstant.ErrInternalServerError.Error()
	if param.Message != nil {
		message = *param.Message
	}

	param.Gin.JSON(param.Code, ApiResponse{
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
	})
}
