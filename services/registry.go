package services

import (
	"hesdastore/api-ppob/clients/config"
	clients "hesdastore/api-ppob/clients/digiflazz"
	"hesdastore/api-ppob/pkg/rabbitmq"
	"hesdastore/api-ppob/repositories"
	serviceAccount "hesdastore/api-ppob/services/account"
	serviceAuth "hesdastore/api-ppob/services/auth"
	serviceBill "hesdastore/api-ppob/services/billing"
	services "hesdastore/api-ppob/services/brand"
	serviceProduct "hesdastore/api-ppob/services/product"
	serviceTransaction "hesdastore/api-ppob/services/transaction"
	serviceWebhook "hesdastore/api-ppob/services/webhook"
)

type IServiceRegistry interface {
	Brand() services.BrandService
	AuthApi() serviceAuth.AuthApiService
	Account() serviceAccount.AccountService
	Product() serviceProduct.ProductService
	Transaction() serviceTransaction.TransactionService
	Billing() serviceBill.BillingService
	Webhook() serviceWebhook.WebhookService
}

type Registry struct {
	brandService       services.BrandService
	authApiService     serviceAuth.AuthApiService
	accountService     serviceAccount.AccountService
	productService     serviceProduct.ProductService
	transactionService serviceTransaction.TransactionService
	billService        serviceBill.BillingService
	webhookService     serviceWebhook.WebhookService
}

func NewServiceRegistry(repository repositories.IRepoRegistry, digifalzz clients.IDigiflazzClient, clientConfig config.IClientConfig, managerRb rabbitmq.IManager) IServiceRegistry {
	return &Registry{
		brandService:       services.NewBrandServiceImpl(repository),
		authApiService:     serviceAuth.NewAuthApiServiceImpl(repository),
		accountService:     serviceAccount.NewAccountServiceImpl(repository),
		productService:     serviceProduct.NewProductServiceImpl(repository),
		transactionService: serviceTransaction.NewTransactionServiceImpl(repository, digifalzz, clientConfig, managerRb),
		billService:        serviceBill.NewBillingServiceImpl(repository, digifalzz, clientConfig),
		webhookService:     serviceWebhook.NewWebhookServiceImpl(repository),
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

func (r *Registry) Product() serviceProduct.ProductService {
	return r.productService
}

func (r *Registry) Transaction() serviceTransaction.TransactionService {
	return r.transactionService
}

func (r *Registry) Billing() serviceBill.BillingService {
	return r.billService
}

func (r *Registry) Webhook() serviceWebhook.WebhookService {
	return r.webhookService
}
