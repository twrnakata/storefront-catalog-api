package model

type CreateProductRequestModel struct {
	Name        string
	Description *string
	SalePrice   *float64
	Price       float64
}

type CreateProductModel struct {
	ID   string
	Name string
}
