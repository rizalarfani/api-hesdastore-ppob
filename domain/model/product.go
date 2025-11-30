package model

type Product struct {
	ProductID   int    `db:"id"`
	Metode      string `db:"metode"`
	ProductCode string `db:"product_code"`
	ProductName string `db:"product_name"`
	Category
	Brand
	Type           string  `db:"type"`
	Hpp            *int    `db:"hpp"`
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
	ID   int    `db:"category_id"`
	Name string `db:"category_name"`
}
