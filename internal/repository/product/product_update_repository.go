package product

import (
	"context"
	"errors"

	"gorm.io/gorm"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
)

var ErrProductNotFound = errors.New("product not found")

type ProductUpdateRepository struct {
	DB *gorm.DB
}

func NewProductUpdateRepository(database *gorm.DB) (*ProductUpdateRepository, error) {
	if database == nil {
		return nil, gorm.ErrInvalidDB
	}
	return &ProductUpdateRepository{DB: database}, nil
}

func (repository *ProductUpdateRepository) UpdateProduct(executionContext context.Context, request *repositorymodel.UpdateProductRequestModel) error {
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
			return ErrProductNotFound
		}
		return err
	}

	return repository.DB.WithContext(executionContext).Model(&ProductRecord{}).Where("id = ?", request.ID).Updates(updates).Error
}
