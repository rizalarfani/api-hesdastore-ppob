package controllers

import (
	accountC "hesdastore/api-ppob/controllers/account"
	controllers "hesdastore/api-ppob/controllers/brand"
	"hesdastore/api-ppob/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	BrandController() controllers.IBrandController
	AccountController() accountC.IAccountController
}

func NewControllerRegistry(s services.IServiceRegistry) IControllerRegistry {
	return &Registry{
		service: s,
	}
}

func (r *Registry) BrandController() controllers.IBrandController {
	return controllers.NewBrandController(r.service)
}

func (r *Registry) AccountController() accountC.IAccountController {
	return accountC.NewAccountController(r.service)
}
