package dto

import "hesdastore/api-ppob/constants"

type TransactionOrderRequest struct {
	ProductCode string `json:"product_code" validate:"required"`
	CustomerNo  string `json:"customer_no" validate:"required"`
	CallbackURL string `json:"callback_url" validate:"omitempty,url"`
}

type TransactionUpdateBalanceRequest struct {
	UserID     int `db:"id"`
	NewBalance int `db:"salod"`
}

type TransactionUpdateRequest struct {
	TrxId     string
	Response  string
	Status    int
	StatusMsg string
	Sn        string
}

type TransactionOrderResponse struct {
	TransactionsID string                            `json:"transaction_id"`
	ProductCode    string                            `json:"product_code"`
	ProductName    string                            `json:"product_name"`
	SN             *string                           `json:"sn"`
	Status         constants.TransactionStatusString `json:"status"`
	Message        string                            `json:"message"`
}

type TransactionHistoryResponse struct {
	TransactionsID string                            `json:"transaction_id"`
	ProductName    string                            `json:"product_name"`
	Brand          *BrandResponse                    `json:"brand"`
	Category       *CategoryResponse                 `json:"category"`
	Price          int                               `json:"price"`
	CustomerNo     string                            `json:"customer_no"`
	SN             *string                           `json:"sn"`
	Status         constants.TransactionStatusString `json:"status"`
	Message        *string                           `json:"message"`
}
