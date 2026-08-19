package product

import (
	"context"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

type CreateProductService struct {
	Repository domainproduct.CreateProductRepository
}

func (service *CreateProductService) Create(executionContext context.Context, request *servicemodel.CreateProductRequestModel, response *servicemodel.CreateProductResponseModel) error {
	if service.Repository == nil {
		return apperror.ErrCreateProductRepositoryNotConfigured
	}

	createdProductModel := &repositorymodel.CreateProductModel{}
	err := service.Repository.CreateProduct(executionContext, &repositorymodel.CreateProductRequestModel{
		Name:        request.Name,
		Description: request.Description,
		SalePrice:   request.SalePrice,
		Price:       *request.Price,
	}, createdProductModel)
	if err != nil {
		return err
	}

	response.ID = createdProductModel.ID
	response.Name = createdProductModel.Name
	return nil
}
