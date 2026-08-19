package model

import "github.com/twrnakata/storefront-catalog-api/pkg/optional"

type UpdateProductRequestModel struct {
	ID          string
	Name        optional.Optional[string]
	Description optional.Optional[string]
	SalePrice   optional.Optional[float64]
	Price       optional.Optional[float64]
}
