package response

import (
	"errors"
	"hesdastore/api-ppob/constants"
	errConstant "hesdastore/api-ppob/constants/error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message any         `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
}

func HttpResponse(param ParamHTTPResp) {
	// SUCCESS
	if param.Err == nil {
		code := param.Code
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
	code := http.StatusInternalServerError
	msg := errConstant.ErrInternalServerError.Error()

	switch {
	case errors.Is(param.Err, errConstant.ErrBadRequest):
		code = http.StatusBadRequest
		msg = param.Err.Error()
	case errors.Is(param.Err, errConstant.ErrProductNotFound):
		code = http.StatusNotFound
		msg = param.Err.Error()
	default:
		logrus.Errorf("error: %v", param.Err)
	}

	if param.Message != nil {
		msg = *param.Message
	}

	param.Gin.JSON(code, ApiResponse{
		Status:  constants.Error,
		Message: msg,
	})
	return
}
