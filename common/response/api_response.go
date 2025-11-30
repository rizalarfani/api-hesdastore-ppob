package response

import (
	"errors"
	"hesdastore/api-ppob/constants"
	errConstant "hesdastore/api-ppob/constants/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
	Errors  interface{}
}

func HttpResponse(param ParamHTTPResp) {
	code := param.Code

	// SUCCESS
	if param.Err == nil {
		if code == 0 {
			code = http.StatusOK
		}

		msg := http.StatusText(code)
		if param.Message != nil {
			msg = *param.Message
		}

		param.Gin.JSON(code, ApiResponse{
			Status:  constants.Success,
			Message: msg,
			Data:    param.Data,
		})
		return
	}

	// ERROR
	message := errConstant.ErrInternalServerError.Error()
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		if errConstant.ErrMapping(param.Err) {
			message = param.Err.Error()
		} else {
			message = param.Err.Error()
		}

		if errors.Is(param.Err, errConstant.ErrInternalServerError) {
			code = http.StatusInternalServerError
		}
	}

	param.Gin.JSON(code, ApiResponse{
		Status:  constants.Error,
		Message: message,
		Errors:  param.Errors,
	})
}
