package dto

type BrandResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
