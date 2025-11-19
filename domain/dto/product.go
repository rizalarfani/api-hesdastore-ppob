package dto

type ProductResponse struct {
	ProductCode string            `json:"product_code"`
	ProductName string            `json:"product_name"`
	Category    *CategoryResponse `json:"category"`
	Brand       *BrandResponse    `json:"brand"`
	Type        string            `json:"type"`
	Price       int               `json:"price,omitempty"`
	Admin       int               `json:"admin,omitempty"`
	Commission  int               `json:"commission,omitempty"`
	Status      string            `json="status"`
	StartCutOff string            `json:"start_cut_off,omitempty"`
	EndCutOff   string            `json:"end_cut_off,omitempty"`
	Description string            `json:"description"`
}

type CategoryResponse struct {
	Name string `json:"name"`
}
