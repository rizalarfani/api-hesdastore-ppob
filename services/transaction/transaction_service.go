package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/domain/model"
)

type TransactionService interface {
	Order(context.Context, *dto.TransactionOrderRequest, *model.ApiUser) (*dto.TransactionOrderResponse, error)
}
