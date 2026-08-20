package handler

import (
	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

type ProductHandler struct {
	ProductService domainproduct.ProductService
}

func NewProductHandler(productService domainproduct.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (handler *ProductHandler) validatePriceNonNegative(price float64) error {
	if price < 0 {
		return apperror.ErrPriceMustBeGreaterThanOrEqualToZero
	}
	return nil
}

func (handler *ProductHandler) validateSalePriceNonNegative(salePrice float64) error {
	if salePrice < 0 {
		return apperror.ErrSalePriceMustBeGreaterThanOrEqualToZero
	}
	return nil
}

func (handler *ProductHandler) validateSalePriceNotExceedPrice(salePrice, price float64) error {
	if salePrice > price {
		return apperror.ErrSalePriceMustNotExceedPrice
	}
	return nil
}
