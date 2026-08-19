package product

import (
	"context"
	"errors"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	repositoryproduct "github.com/twrnakata/storefront-catalog-api/internal/repository/product"
	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

type UpdateProductService struct {
	Repository domainproduct.UpdateProductRepository
}

func (service *UpdateProductService) Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error {
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
		if errors.Is(err, repositoryproduct.ErrProductNotFound) {
			return domainproduct.ErrProductNotFound
		}
		return err
	}
	return nil
}
