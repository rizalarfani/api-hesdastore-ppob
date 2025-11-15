package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type AuhtApiRepository interface {
	FindByUsername(ctx context.Context, username string) (*model.ApiUser, error)
}
