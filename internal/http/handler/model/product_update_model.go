package model

import "github.com/twrnakata/storefront-catalog-api/pkg/optional"

type UpdateProductRequestModel struct {
	ID          string                     `json:"-"`
	Name        optional.Optional[string]  `json:"name"`
	Description optional.Optional[string]  `json:"description"`
	SalePrice   optional.Optional[float64] `json:"sale_price"`
	Price       optional.Optional[float64] `json:"price"`
}
