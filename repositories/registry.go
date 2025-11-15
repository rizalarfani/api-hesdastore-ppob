package repositories

import (
	repositoriesAccount "hesdastore/api-ppob/repositories/account"
	repositoriesAuth "hesdastore/api-ppob/repositories/auth"
	repositories "hesdastore/api-ppob/repositories/brand"

	"github.com/jmoiron/sqlx"
)

type IRepoRegistry interface {
	Brand() repositories.BrandRepository
	AuthApi() repositoriesAuth.AuhtApiRepository
	Account() repositoriesAccount.AccountRepository
}

type Registry struct {
	brandRepo   repositories.BrandRepository
	authApiRepo repositoriesAuth.AuhtApiRepository
	accountRepo repositoriesAccount.AccountRepository
}

func NewRepositoryRegistry(db *sqlx.DB) IRepoRegistry {
	return &Registry{
		brandRepo:   repositories.NewBrandRepositoryImpl(db),
		authApiRepo: repositoriesAuth.NewAuthApiRepositoryImpl(db),
		accountRepo: repositoriesAccount.NewAccountRepositoryImpl(db),
	}
}

func (r *Registry) Brand() repositories.BrandRepository {
	return r.brandRepo
}

func (r *Registry) AuthApi() repositoriesAuth.AuhtApiRepository {
	return r.authApiRepo
}

func (r *Registry) Account() repositoriesAccount.AccountRepository {
	return r.accountRepo
}
