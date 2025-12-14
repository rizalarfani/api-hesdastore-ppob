package controllers

import (
	accountC "hesdastore/api-ppob/controllers/account"
	billingC "hesdastore/api-ppob/controllers/billing"
	controllers "hesdastore/api-ppob/controllers/brand"
	productC "hesdastore/api-ppob/controllers/product"
	transactionC "hesdastore/api-ppob/controllers/transaction"
	webhookC "hesdastore/api-ppob/controllers/webhook"
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
	BillingController() billingC.IBillingController
	WebhookController() webhookC.IWebhookController
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

func (r *Registry) BillingController() billingC.IBillingController {
	return billingC.NewBillingController(r.service)
}

func (r *Registry) WebhookController() webhookC.IWebhookController {
	return webhookC.NewWebhookController(r.service)
}
