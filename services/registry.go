package services

import (
	"hesdastore/api-ppob/repositories"
	serviceAccount "hesdastore/api-ppob/services/account"
	serviceAuth "hesdastore/api-ppob/services/auth"
	services "hesdastore/api-ppob/services/brand"
)

type IServiceRegistry interface {
	Brand() services.BrandService
	AuthApi() serviceAuth.AuthApiService
	Account() serviceAccount.AccountService
}

type Registry struct {
	brandService   services.BrandService
	authApiService serviceAuth.AuthApiService
	accountService serviceAccount.AccountService
}

func NewServiceRegistry(repository repositories.IRepoRegistry) IServiceRegistry {
	return &Registry{
		brandService:   services.NewBrandServiceImpl(repository),
		authApiService: serviceAuth.NewAuthApiServiceImpl(repository),
		accountService: serviceAccount.NewAccountServiceImpl(repository),
	}
}

func (r *Registry) Brand() services.BrandService {
	return r.brandService
}

func (r *Registry) AuthApi() serviceAuth.AuthApiService {
	return r.authApiService
}

func (r *Registry) Account() serviceAccount.AccountService {
	return r.accountService
}
