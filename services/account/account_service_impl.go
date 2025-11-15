package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/repositories"
)

type AccountServiceImpl struct {
	repository repositories.IRepoRegistry
}

func NewAccountServiceImpl(repo repositories.IRepoRegistry) AccountService {
	return &AccountServiceImpl{
		repository: repo,
	}
}

func (s *AccountServiceImpl) FindBalanceUser(ctx context.Context, username string) (*dto.AccountResponse, error) {
	data, err := s.repository.Account().FindBalanceUser(ctx, username)
	if err != nil {
		return nil, err
	}

	account := dto.AccountResponse{
		Name:    data.Name,
		Balance: data.Balance,
	}
	return &account, nil
}
