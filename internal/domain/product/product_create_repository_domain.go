package product

import (
	"context"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
)

type CreateProductRepository interface {
	CreateProduct(executionContext context.Context, request *repositorymodel.CreateProductRequestModel, response *repositorymodel.CreateProductModel) error
}
