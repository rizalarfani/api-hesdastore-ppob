package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type BrandRepository interface {
	FindAll(ctx context.Context) ([]*model.Brand, error)
}
