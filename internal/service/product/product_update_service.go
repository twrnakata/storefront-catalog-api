package product

import (
	"context"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

func (service *ProductService) Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error {
	if service.Repository == nil {
		return apperror.ErrUpdateProductRepositoryNotConfigured
	}

	err := service.Repository.UpdateProduct(executionContext, &repositorymodel.UpdateProductRequestModel{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		SalePrice:   request.SalePrice,
		Price:       request.Price,
	})
	if err != nil {
		return err
	}
	return nil
}
