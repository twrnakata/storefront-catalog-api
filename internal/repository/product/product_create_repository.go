package product

import (
	"context"

	"github.com/google/uuid"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
)

func (repository *ProductRepository) CreateProduct(executionContext context.Context, request *repositorymodel.CreateProductRequestModel, response *repositorymodel.CreateProductModel) error {
	record := ProductRecord{
		ID:          uuid.NewString(),
		Name:        request.Name,
		Description: request.Description,
		SalePrice:   request.SalePrice,
		Price:       request.Price,
	}
	if err := repository.DB.WithContext(executionContext).Create(&record).Error; err != nil {
		return err
	}
	response.ID = record.ID
	response.Name = record.Name
	return nil
}
