package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
)

type AuthApiService interface {
	ValidateAuthApi(ctx context.Context, request *dto.ApiUserRequest) (*model.ApiUser, error)
}
