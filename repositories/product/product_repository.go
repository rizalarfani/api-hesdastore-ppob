package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type ProductRepository interface {
	FindAllPrabayar(ctx context.Context) ([]*model.Product, error)
	FindAllPascabayar(ctx context.Context) ([]*model.Product, error)
	FindByProductCode(ctx context.Context, productCode string) (*model.Product, error)
}
