package apperror

import "errors"

var (
	ErrCreateProductServiceNotInitialized   = errors.New("create product service not initialized")
	ErrCreateProductRepositoryNotConfigured = errors.New("create product repository not configured")
	ErrUpdateProductServiceNotInitialized   = errors.New("update product service not initialized")
	ErrUpdateProductRepositoryNotConfigured = errors.New("update product repository not configured")
	ErrInvalidJSONBody                      = errors.New("invalid json body")
	ErrNameRequired                         = errors.New("name is required")
	ErrPriceRequired                        = errors.New("price is required")
	ErrDatabaseURLRequired                  = errors.New("DATABASE_URL is required")
	ErrInvalidProductID                     = errors.New("invalid product id")
	ErrAtLeastOneFieldRequired              = errors.New("at least one field is required")
	ErrNameCannotBeNull                     = errors.New("name cannot be null")
	ErrPriceCannotBeNull                    = errors.New("price cannot be null")
	ErrPriceMustBeGreaterThanOrEqualToZero  = errors.New("price must be greater than or equal to 0")
	ErrSalePriceMustBeGreaterThanOrEqualToZero = errors.New("sale_price must be greater than or equal to 0")
	ErrSalePriceMustNotExceedPrice          = errors.New("sale_price must not exceed price")
)
