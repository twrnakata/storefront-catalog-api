package product

import (
	"context"
	"errors"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	CreateProduct(executionContext context.Context, request *repositorymodel.CreateProductRequestModel, response *repositorymodel.CreateProductModel) error
	UpdateProduct(executionContext context.Context, request *repositorymodel.UpdateProductRequestModel) error
}

type ProductService interface {
	Create(executionContext context.Context, request *servicemodel.CreateProductRequestModel, response *servicemodel.CreateProductResponseModel) error
	Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error
}
