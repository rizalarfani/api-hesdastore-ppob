package dto

import "hesdastore/api-ppob/constants"

type PayBillRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
	ProductCode   string `json:"product_code" validate:"required"`
	CustomerNo    string `json:"customer_no" validate:"required"`
	CallbackURL   string `json:"callback_url" validate:"omitempty,url"`
}

type PayBillResponse struct {
	Data DataPayBill `json:"data"`
}

type DataPayBill struct {
	TransactionsID string                            `json:"transaction_id"`
	ProductCode    string                            `json:"product_code"`
	Brand          *BrandResponse                    `json:"brand"`
	Category       *CategoryResponse                 `json:"category"`
	CustomerNo     string                            `json:"customer_no"`
	CustomerName   string                            `json:"customer_name"`
	Price          int                               `json:"price"`
	Sn             string                            `json:"sn"`
	Status         constants.TransactionStatusString `json:"status"`
	Message        string                            `json:"message"`
}

type InquiryBillRequest struct {
	ProductCode string `json:"product_code" validate:"required"`
	CustomerNo  string `json:"customer_no" validate:"required"`
}

type InquiryBillingResponse struct {
	TransactionsID string                            `json:"transaction_id"`
	ProductCode    string                            `json:"product_code"`
	Brand          *BrandResponse                    `json:"brand"`
	Category       *CategoryResponse                 `json:"category"`
	CustomerName   string                            `json:"customer_name"`
	Price          int                               `json:"price"`
	Status         constants.TransactionStatusString `json:"status"`
	Message        string                            `json:"message"`
	Detail         any                               `json:"description"`
}

type InternetDescResponse struct {
	BillingSheet int                      `json:"bill_sheet"`
	Detail       []InternetDetailResponse `json:"details"`
}

type InternetDetailResponse struct {
	Period     string `json:"period"`
	BillValue  string `json:"bill_value"`
	PriceAdmin string `json:"price_admin"`
}

type InternetDescRespDigiflazz struct {
	BillingSheet int                           `json:"lembar_tagihan"`
	Detail       []InternetDetailRestDigiflazz `json:"detail"`
}

type InternetDetailRestDigiflazz struct {
	Period     string `json:"periode"`
	BillValue  string `json:"nilai_tagihan"`
	PriceAdmin string `json:"admin"`
}

type PlnDescResponse struct {
	Tarif        string              `json:"tarif"`
	Power        int                 `json:"power"`
	BillingSheet int                 `json:"bill_sheet"`
	Detail       []PlnDetailResponse `json:"details"`
}

type PlnDetailResponse struct {
	Period     string `json:"period"`
	BillValue  string `json:"bill_value"`
	PriceAdmin string `json:"price_admin"`
	Fine       string `json:"fine"`
}

type PlnDescRespDigiflazz struct {
	Tarif        string                   `json:"tarif"`
	Power        int                      `json:"daya"`
	BillingSheet int                      `json:"lembar_tagihan"`
	Detail       []PlnDetailRespDigiflazz `json:"detail"`
}

type PlnDetailRespDigiflazz struct {
	Period     string `json:"periode"`
	BillValue  string `json:"nilai_tagihan"`
	PriceAdmin string `json:"admin"`
	Fine       string `json:"denda"`
}

type PdamDescResponse struct {
	Tarif        string               `json:"tarif"`
	BillingSheet int                  `json:"bill_sheet"`
	Address      string               `json:"address"`
	DueDate      string               `json:"due_date"`
	Detail       []PdamDetailResponse `json:"details"`
}

type PdamDetailResponse struct {
	Period    string `json:"period"`
	BillValue string `json:"bill_value"`
	Fine      string `json:"fine"`
}

type PdamDescRespDigiflazz struct {
	Tarif        string                    `json:"tarif"`
	BillingSheet int                       `json:"lembar_tagihan"`
	Address      string                    `json:"alamat"`
	DueDate      string                    `json:"jatuh_tempo"`
	Detail       []PdamDetailRespDigiflazz `json:"detail"`
}

type PdamDetailRespDigiflazz struct {
	Period    string `json:"periode"`
	BillValue string `json:"nilai_tagihan"`
	Fine      string `json:"denda"`
}

type BpjsKesDescResponse struct {
	BillingSheet       int                     `json:"bill_sheet"`
	Address            string                  `json:"address"`
	NumberParticipants string                  `json:"number_participants"`
	Detail             []BpjsKesDetailResponse `json:"details"`
}

type BpjsKesDetailResponse struct {
	Period string `json:"period"`
}

type BpjsKesDescRespDigiflazz struct {
	BillingSheet       int                          `json:"lembar_tagihan"`
	Address            string                       `json:"alamat"`
	NumberParticipants string                       `json:"jumlah_partisipan"`
	Detail             []BpjsKesDetailRespDigiflazz `json:"detail"`
}

type BpjsKesDetailRespDigiflazz struct {
	Period string `json:"periode"`
}
