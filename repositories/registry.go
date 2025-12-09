package repositories

import (
	repositoriesAccount "hesdastore/api-ppob/repositories/account"
	repositoriesAuth "hesdastore/api-ppob/repositories/auth"
	repositoriesBilling "hesdastore/api-ppob/repositories/billing"
	repositories "hesdastore/api-ppob/repositories/brand"
	repositoriesConfig "hesdastore/api-ppob/repositories/config"
	repositoriesProduct "hesdastore/api-ppob/repositories/product"
	repositoriesTransaction "hesdastore/api-ppob/repositories/transaction"

	"github.com/jmoiron/sqlx"
)

type IRepoRegistry interface {
	Brand() repositories.BrandRepository
	AuthApi() repositoriesAuth.AuhtApiRepository
	Account() repositoriesAccount.AccountRepository
	Product() repositoriesProduct.ProductRepository
	Transaction() repositoriesTransaction.TransactionRepository
	Config() repositoriesConfig.ConfigRepository
	GetTx() *sqlx.DB
	Billing() repositoriesBilling.BillingRepository
}

type Registry struct {
	db *sqlx.DB
}

func NewRepositoryRegistry(db *sqlx.DB) IRepoRegistry {
	return &Registry{
		db: db,
	}
}

func (r *Registry) Brand() repositories.BrandRepository {
	return repositories.NewBrandRepositoryImpl(r.db)
}

func (r *Registry) AuthApi() repositoriesAuth.AuhtApiRepository {
	return repositoriesAuth.NewAuthApiRepositoryImpl(r.db)
}

func (r *Registry) Account() repositoriesAccount.AccountRepository {
	return repositoriesAccount.NewAccountRepositoryImpl(r.db)
}

func (r *Registry) Product() repositoriesProduct.ProductRepository {
	return repositoriesProduct.NewProductRepositoryImpl(r.db)
}

func (r *Registry) Transaction() repositoriesTransaction.TransactionRepository {
	return repositoriesTransaction.NewProductRepositoryImpl(r.db)
}

func (r *Registry) Config() repositoriesConfig.ConfigRepository {
	return repositoriesConfig.NewConfigRepositoryImpl(r.db)
}

func (r *Registry) GetTx() *sqlx.DB {
	return r.db
}

func (r *Registry) Billing() repositoriesBilling.BillingRepository {
	return repositoriesBilling.NewBillingRepositoryImpl(r.db)
}
