package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	CreateOrder(ctx context.Context, tx *sqlx.Tx, order *model.TransactionOrder) (*model.TransactionOrder, error)
	UpdateBalance(ctx context.Context, tx *sqlx.Tx, req *dto.TransactionUpdateBalanceRequest) error
	GetTransactionByRefID(ctx context.Context, refId string) (*model.TransactionOrder, error)
	UpdateTransaction(ctx context.Context, tx *sqlx.Tx, req *dto.TransactionUpdateRequest) error
}
