package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type AccountRepository interface {
	FindBalanceUser(ctx context.Context, username string) (*model.Account, error)
	FindBalanceUserByUserId(ctx context.Context, userId int) (*model.Account, error)
}
