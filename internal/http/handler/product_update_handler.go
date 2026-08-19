package handler

import (
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	handlermodel "github.com/twrnakata/storefront-catalog-api/internal/http/handler/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
	"github.com/twrnakata/storefront-catalog-api/pkg/caller"
	"github.com/twrnakata/storefront-catalog-api/pkg/optional"
)

type ProductUpdateHandler struct {
	UpdateService domainproduct.UpdateProductService
}

func (handler *ProductUpdateHandler) Update(c *fiber.Ctx) error {
	if handler.UpdateService == nil {
		return caller.InternalServerError(c, apperror.ErrUpdateProductServiceNotInitialized)
	}

	var request handlermodel.UpdateProductRequestModel
	err := validateUpdateProductRequest(c, &request)
	if err != nil {
		return caller.BadRequest(c, err.Error())
	}

	err = handler.UpdateService.Update(context.Background(), &servicemodel.UpdateProductRequestModel{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		SalePrice:   request.SalePrice,
		Price:       request.Price,
	})
	if err != nil {
		if errors.Is(err, domainproduct.ErrProductNotFound) {
			return caller.NotFound(c, err.Error())
		}
		return caller.InternalServerError(c, err)
	}

	return caller.Success(c, nil)
}

func validateUpdateProductRequest(c *fiber.Ctx, request *handlermodel.UpdateProductRequestModel) error {
	request.ID = strings.TrimSpace(c.Params("id"))
	if request.ID == "" {
		return apperror.ErrInvalidProductID
	}
	if _, err := uuid.Parse(request.ID); err != nil {
		return apperror.ErrInvalidProductID
	}

	if err := c.BodyParser(request); err != nil {
		return apperror.ErrInvalidJSONBody
	}

	if request.Name.IsSet() {
		if request.Name.IsNull() {
			return apperror.ErrNameCannotBeNull
		}
		trimmedName := strings.TrimSpace(request.Name.Value())
		if trimmedName == "" {
			return apperror.ErrNameRequired
		}
		request.Name = optional.From(trimmedName)
	}
	if request.Price.IsSet() && request.Price.IsNull() {
		return apperror.ErrPriceCannotBeNull
	}

	if !request.Name.IsSet() && !request.Description.IsSet() && !request.SalePrice.IsSet() && !request.Price.IsSet() {
		return apperror.ErrAtLeastOneFieldRequired
	}
	return nil
}
