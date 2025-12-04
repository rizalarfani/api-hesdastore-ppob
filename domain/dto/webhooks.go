package dto

type DigifalzzWebhooksPayload struct {
	Data DataPayload `json:"data"`
}

type DataPayload struct {
	RefID      string  `json:"ref_id"`
	CustomerNo string  `json:"customer_no"`
	SKUCode    string  `json:"buyer_sku_code"`
	Message    string  `json:"message"`
	Status     string  `json:"status"`
	Rc         string  `json:"rc"`
	Sn         *string `json:"sn"`
	LastSaldo  int     `json:"buyer_last_saldo"`
	Price      int     `json:"price"`
	Telegram   string  `json:"tele"`
	Wahtsapp   string  `json:"wa"`
}
