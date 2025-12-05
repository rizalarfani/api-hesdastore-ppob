package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
)

type TransactionService interface {
	GetHistory(ctx context.Context, trxID string, userId int) ([]*dto.TransactionHistoryResponse, error)
	Order(context.Context, *dto.TransactionOrderRequest, *model.ApiUser) (*dto.TransactionOrderResponse, error)
	Webhooks(ctx context.Context, headerHubSignature string, rawBody []byte) error
}
