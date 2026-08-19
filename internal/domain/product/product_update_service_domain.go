package product

import (
	"context"

	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
)

type UpdateProductService interface {
	Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error
}
