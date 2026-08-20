package product

import (
	"context"
	"errors"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	"gorm.io/gorm"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
)

func (repository *ProductRepository) UpdateProduct(executionContext context.Context, request *repositorymodel.UpdateProductRequestModel) error {
	updates := map[string]any{}
	if request.Name.IsSet() {
		updates["name"] = request.Name.Value()
	}
	if request.Description.IsSet() {
		if request.Description.IsNull() {
			updates["description"] = nil
		} else {
			updates["description"] = request.Description.Value()
		}
	}
	if request.SalePrice.IsSet() {
		if request.SalePrice.IsNull() {
			updates["sale_price"] = nil
		} else {
			updates["sale_price"] = request.SalePrice.Value()
		}
	}
	if request.Price.IsSet() {
		updates["price"] = request.Price.Value()
	}

	var existing ProductRecord
	err := repository.DB.WithContext(executionContext).Where("id = ?", request.ID).First(&existing).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domainproduct.ErrProductNotFound
		}
		return err
	}

	return repository.DB.WithContext(executionContext).Model(&ProductRecord{}).Where("id = ?", request.ID).Updates(updates).Error
}
