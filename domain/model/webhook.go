package model

type Webhook struct {
	EventType      string `db:"event_type"`
	DeleveryRef    string `db:"delivery_ref"`
	Endpoint       string `db:"endpoint"`
	RequestBody    string `db:"request_body"`
	ResponseBody   string `db:"response_body"`
	ResponseStatus int    `db:"response_status"`
	ResponseError  string `db:"response_error"`
	Signature      string `db:"signature"`
}
