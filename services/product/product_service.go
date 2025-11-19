package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
)

type ProductService interface {
	FindAllPrabayar(ctx context.Context, role int) ([]*dto.ProductResponse, error)
	FindAllPascabayar(ctx context.Context, role int) ([]*dto.ProductResponse, error)
	FindByProductCode(ctx context.Context, productCode string, role int) (*dto.ProductResponse, error)
}
