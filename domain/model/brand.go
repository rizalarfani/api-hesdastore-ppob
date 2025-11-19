package model

type Brand struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
