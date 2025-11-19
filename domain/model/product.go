package model

type Product struct {
	Metode      string `db:"metode"`
	ProductCode string `db:"product_code"`
	ProductName string `db:"product_name"`
	Category
	Brand
	Type           string  `db:"type"`
	PriceSeller    *int    `db:"price_seller"`
	PriceReseller  *int    `db:"price_reseller"`
	Admin          *int    `db:"admin"`
	Commission     *int    `db:"commission"`
	SellerStatus   string  `db:"seller_status"`
	ResellerStatus string  `db:"reseller_status"`
	StartCutOff    *string `db:"start_cut_off"`
	EndCutOff      *string `db:"end_cut_off"`
	Description    string  `db:"description"`
}

type Category struct {
	Name string `db:"category_name"`
}
