package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
)

type AccountService interface {
	FindBalanceUser(ctx context.Context, username string) (*dto.AccountResponse, error)
}
