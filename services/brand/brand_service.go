package services

import (
	"context"
	"hesdastore/api-ppob/domain/dto"
)

type BrandService interface {
	FindAll(ctx context.Context) ([]*dto.BrandResponse, error)
}
