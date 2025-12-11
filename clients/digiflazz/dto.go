package clients

import "encoding/json"

type TopupRequest struct {
	Username   string `json:"username"`
	SKUCode    string `json:"buyer_sku_code"`
	CustomerNo string `json:"customer_no"`
	RefID      string `json:"ref_id"`
	Signature  string `json:"sign"`
	CalbackURL string `json:"cb_url"`
}

type TopupResponse struct {
	Data DataTopup `json:"data"`
}

type DataTopup struct {
	RefID      string  `json:"ref_id"`
	CustomerNo string  `json:"customer_no"`
	SKUCode    string  `json:"buyer_sku_code"`
	Message    string  `json:"message"`
	Status     string  `json:"status"`
	Rc         string  `json:"rc"`
	Sn         *string `json:"sn"`
}

type BillPayRequest struct {
	Commands   string `json:"commands"`
	Username   string `json:"username"`
	SKUCode    string `json:"buyer_sku_code"`
	CustomerNo string `json:"customer_no"`
	RefID      string `json:"ref_id"`
	Signature  string `json:"sign"`
}

type BillPaymentResponse struct {
	Data DataBillPayment `json:"data"`
}

type DataBillPayment struct {
	RefID         string  `json:"ref_id"`
	CustomerNo    string  `json:"customer_no"`
	CustomerName  string  `json:"customer_name"`
	SKUCode       string  `json:"buyer_sku_code"`
	OriginalPrice int     `json:"price"`
	Price         int     `json:"selling_price"`
	Message       string  `json:"message"`
	Status        string  `json:"status"`
	Rc            string  `json:"rc"`
	Sn            *string `json:"sn"`
}

type InquiryRequest struct {
	Command    string `json:"commands"`
	Username   string `json:"username"`
	SKUCode    string `json:"buyer_sku_code"`
	CustomerNo string `json:"customer_no"`
	RefID      string `json:"ref_id"`
	Signature  string `json:"sign"`
}

type InquiryResponse struct {
	Data DataInquiry `json:"data"`
}

type DataInquiry struct {
	RefID         string          `json:"ref_id"`
	CustomerNo    string          `json:"customer_no"`
	CustomerName  string          `json:"customer_name"`
	SKUCode       string          `json:"buyer_sku_code"`
	OriginalPrice int             `json:"price"`
	Price         int             `json:"selling_price"`
	Message       string          `json:"message"`
	Status        string          `json:"status"`
	Rc            string          `json:"rc"`
	Desc          json.RawMessage `json:"desc"`
}
