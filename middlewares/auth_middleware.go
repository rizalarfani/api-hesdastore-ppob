package middlewares

import (
	"encoding/base64"
	"hesdastore/api-ppob/common/response"
	"hesdastore/api-ppob/constants"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/services"
	"net/http"
	"strings"

	errConstant "hesdastore/api-ppob/constants/error"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service services.IServiceRegistry
}

func NewAuthMiddleware(authService services.IServiceRegistry) *AuthMiddleware {
	return &AuthMiddleware{service: authService}
}

func validateBasicAuthAndApiKey(c *gin.Context, m *AuthMiddleware) (int, error) {
	authBasic := c.GetHeader(constants.Authorization)
	if !strings.HasPrefix(authBasic, "Basic ") {
		return http.StatusUnauthorized, errConstant.ErrUnauthorized
	}

	encoded := strings.TrimPrefix(authBasic, "Basic ")
	decodedByte, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return http.StatusUnauthorized, errConstant.ErrUnauthorized
	}

	parts := strings.SplitN(string(decodedByte), ":", 2)
	if len(parts) != 2 {
		return http.StatusUnauthorized, errConstant.ErrUnauthorized
	}

	username := parts[0]
	password := parts[1]

	var apiKey string = c.GetHeader(constants.XApiKey)
	if apiKey == "" {
		return http.StatusUnauthorized, errConstant.ErrApiKeyInvalid
	}

	request := &dto.ApiUserRequest{
		Username: username,
		Password: password,
		ApiKey:   apiKey,
	}

	user, err := m.service.AuthApi().ValidateAuthApi(c, request)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusUnauthorized, errConstant.ErrUnauthorized
	}

	c.Set("authUser", user)
	return http.StatusOK, nil
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var statusCode int

		statusCode, err = validateBasicAuthAndApiKey(ctx, m)
		if err != nil {
			ctx.JSON(statusCode, response.ApiResponse{
				Status:  constants.Error,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
