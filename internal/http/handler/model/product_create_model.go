package model

type CreateProductRequestModel struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	SalePrice   *float64 `json:"sale_price"`
	Price       *float64 `json:"price"`
}

type CreateProductResponseModel struct {
	// data1, data2 ชื่อตาม spec ของ API
	Data1 string `json:"data1"` // id สินค้า
	Data2 string `json:"data2"` // ชื่อสินค้า
}
