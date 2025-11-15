package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
	"hesdastore/api-ppob/repositories"

	errConstant "hesdastore/api-ppob/constants/error"

	"golang.org/x/crypto/bcrypt"
)

type AuthApiServiceImpl struct {
	repository repositories.IRepoRegistry
}

func NewAuthApiServiceImpl(repo repositories.IRepoRegistry) AuthApiService {
	return &AuthApiServiceImpl{
		repository: repo,
	}
}

func (s *AuthApiServiceImpl) ValidateAuthApi(ctx context.Context, request *dto.ApiUserRequest) (*model.ApiUser, error) {
	user, err := s.repository.AuthApi().FindByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errConstant.ErrPasswordInCorrect
	}

	if user.APIKey != request.ApiKey {
		return nil, errConstant.ErrApiKeyInvalid
	}

	return user, nil
}
