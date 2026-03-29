package repositories

import (
	"context"
	"hesdastore/api-ppob/domain/model"
)

type AuhtApiRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.ApiUser, error)
	FindSecretKeyByUserID(ctx context.Context, userId int) (string, error)
}
