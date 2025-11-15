package repositories

import (
	repositoriesAuth "hesdastore/api-ppob/repositories/auth"
	repositories "hesdastore/api-ppob/repositories/brand"

	"github.com/jmoiron/sqlx"
)

type IRepoRegistry interface {
	Brand() repositories.BrandRepository
	AuthApi() repositoriesAuth.AuhtApiRepository
}

type Registry struct {
	brandRepo   repositories.BrandRepository
	authApiRepo repositoriesAuth.AuhtApiRepository
}

func NewRepositoryRegistry(db *sqlx.DB) IRepoRegistry {
	return &Registry{
		brandRepo:   repositories.NewBrandRepositoryImpl(db),
		authApiRepo: repositoriesAuth.NewAuthApiRepositoryImpl(db),
	}
}

func (r *Registry) Brand() repositories.BrandRepository {
	return r.brandRepo
}

func (r *Registry) AuthApi() repositoriesAuth.AuhtApiRepository {
	return r.authApiRepo
}
