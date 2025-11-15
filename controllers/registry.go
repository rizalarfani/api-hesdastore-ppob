package controllers

import (
	controllers "hesdastore/api-ppob/controllers/brand"
	"hesdastore/api-ppob/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	BrandController() controllers.IBrandController
}

func NewControllerRegistry(s services.IServiceRegistry) IControllerRegistry {
	return &Registry{
		service: s,
	}
}

func (r *Registry) BrandController() controllers.IBrandController {
	return controllers.NewBrandController(r.service)
}
