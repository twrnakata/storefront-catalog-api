package apperror

import "errors"

var (
	ErrCreateProductServiceNotInitialized   = errors.New("create product service not initialized")
	ErrCreateProductRepositoryNotConfigured = errors.New("create product repository not configured")
	ErrInvalidJSONBody                      = errors.New("invalid json body")
	ErrNameRequired                         = errors.New("name is required")
	ErrPriceRequired                        = errors.New("price is required")
	ErrDatabaseURLRequired                  = errors.New("DATABASE_URL is required")
)
