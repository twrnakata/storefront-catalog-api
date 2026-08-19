package product

import (
	"context"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
)

type UpdateProductRepository interface {
	UpdateProduct(executionContext context.Context, request *repositorymodel.UpdateProductRequestModel) error
}
