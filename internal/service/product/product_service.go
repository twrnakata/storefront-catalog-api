package product

import (
	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
)

type ProductService struct {
	Repository domainproduct.ProductRepository
}

func NewProductService(repository domainproduct.ProductRepository) *ProductService {
	return &ProductService{Repository: repository}
}
