package product

import (
	"context"

	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
)

type CreateProductService interface {
	Create(executionContext context.Context, request *servicemodel.CreateProductRequestModel, response *servicemodel.CreateProductResponseModel) error
}
