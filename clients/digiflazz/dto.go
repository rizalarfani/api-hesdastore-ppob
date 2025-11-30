package clients

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
