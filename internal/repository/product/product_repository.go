package product

import "gorm.io/gorm"

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(database *gorm.DB) (*ProductRepository, error) {
	if database == nil {
		return nil, gorm.ErrInvalidDB
	}
	if err := database.AutoMigrate(&ProductRecord{}); err != nil {
		return nil, err
	}
	return &ProductRepository{DB: database}, nil
}
