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

type TransactionOrderResponse struct {
	TransactionsID string                            `json:"transaction_id"`
	ProductCode    string                            `json:"product_code"`
	ProductName    string                            `json:"product_name"`
	Status         constants.TransactionStatusString `json:"status"`
	Message        string                            `json:"message"`
}
