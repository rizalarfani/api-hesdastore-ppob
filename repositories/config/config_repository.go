package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type ConfigRepository interface {
	GetConfigDigiflazz(ctx context.Context) (*model.Digiflazz, error)
}
