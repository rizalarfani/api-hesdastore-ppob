package controllers

import (
	accountC "hesdastore/api-ppob/controllers/account"
	controllers "hesdastore/api-ppob/controllers/brand"
	productC "hesdastore/api-ppob/controllers/product"
	transactionC "hesdastore/api-ppob/controllers/transaction"
	"hesdastore/api-ppob/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	BrandController() controllers.IBrandController
	AccountController() accountC.IAccountController
	ProductController() productC.IProductController
	TransactionController() transactionC.ITransactionController
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

func (r *Registry) ProductController() productC.IProductController {
	return productC.NewBrandController(r.service)
}

func (r *Registry) TransactionController() transactionC.ITransactionController {
	return transactionC.NewBrandController(r.service)
}
