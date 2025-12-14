package dto

type TransactionUpdateEventWebhook struct {
	TransactionID string `json:"transaction_id"`
	ProductName   string `json:"product_name"`
	CustomerNo    string `json:"customer_no"`
	Price         int    `json:"price"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	SN            string `json:"sn,omitempty"`
	CallbackURL   string `json:"callback_url"`
	Signature     string `json:"signature"`
	EventType     string `json:"event_type"`
}

type WebhooksPayloadToClient struct {
	TransactionID string `json:"transaction_id"`
	ProductName   string `json:"product_name"`
	CustomerNo    string `json:"customer_no"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	SN            string `json:"sn,omitempty"`
	Timestamp     string `json:"timestamp"`
}
