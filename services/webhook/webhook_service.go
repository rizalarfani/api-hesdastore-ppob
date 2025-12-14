package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
)

type WebhookService interface {
	SendWebhook(ctx context.Context, payload *dto.TransactionUpdateEventWebhook) error
}
