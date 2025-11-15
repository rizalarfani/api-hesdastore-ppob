package services

import (
	"hesdastore/api-ppob/repositories"
	serviceAuth "hesdastore/api-ppob/services/auth"
	services "hesdastore/api-ppob/services/brand"
)

type IServiceRegistry interface {
	Brand() services.BrandService
	AuthApi() serviceAuth.AuthApiService
}

type Registry struct {
	brandService   services.BrandService
	authApiService serviceAuth.AuthApiService
}

func NewServiceRegistry(repository repositories.IRepoRegistry) IServiceRegistry {
	return &Registry{
		brandService:   services.NewBrandServiceImpl(repository),
		authApiService: serviceAuth.NewAuthApiServiceImpl(repository),
	}
}

func (r *Registry) Brand() services.BrandService {
	return r.brandService
}

func (r *Registry) AuthApi() serviceAuth.AuthApiService {
	return r.authApiService
}
