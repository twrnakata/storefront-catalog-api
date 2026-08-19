package product

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
)

type ProductCreateRepository struct {
	DB *gorm.DB
}

func NewProductCreateRepository(database *gorm.DB) (*ProductCreateRepository, error) {
	if database == nil {
		return nil, gorm.ErrInvalidDB
	}
	if err := database.AutoMigrate(&ProductRecord{}); err != nil {
		return nil, err
	}
	return &ProductCreateRepository{DB: database}, nil
}

func (repository *ProductCreateRepository) CreateProduct(executionContext context.Context, request *repositorymodel.CreateProductRequestModel, response *repositorymodel.CreateProductModel) error {
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
