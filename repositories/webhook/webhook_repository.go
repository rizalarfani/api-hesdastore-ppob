package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type WebhookRepository interface {
	Create(context.Context, *model.Webhook) error
}
