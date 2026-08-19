package handler

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	handlermodel "github.com/twrnakata/storefront-catalog-api/internal/http/handler/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
	"github.com/twrnakata/storefront-catalog-api/pkg/caller"
)

type ProductCreateHandler struct {
	CreateService domainproduct.CreateProductService
}

func (handler *ProductCreateHandler) Create(c *fiber.Ctx) error {
	if handler.CreateService == nil {
		return caller.InternalServerError(c, apperror.ErrCreateProductServiceNotInitialized)
	}

	var request handlermodel.CreateProductRequestModel
	err := validateCreateProductRequest(c, &request)
	if err != nil {
		return caller.BadRequest(c, err.Error())
	}

	var response servicemodel.CreateProductResponseModel
	err = handler.CreateService.Create(context.Background(), &servicemodel.CreateProductRequestModel{
		Name:        request.Name,
		Description: request.Description,
		SalePrice:   request.SalePrice,
		Price:       request.Price,
	}, &response)
	if err != nil {
		return caller.InternalServerError(c, err)
	}

	responseBody := handlermodel.CreateProductResponseModel{
		Data1: response.ID,   // data1 = id
		Data2: response.Name, // data2 = ชื่อสินค้า
	}
	return caller.Success(c, responseBody)
}

func validateCreateProductRequest(c *fiber.Ctx, request *handlermodel.CreateProductRequestModel) error {
	if err := c.BodyParser(request); err != nil {
		return apperror.ErrInvalidJSONBody
	}

	request.Name = strings.TrimSpace(request.Name)
	if request.Name == "" {
		return apperror.ErrNameRequired
	}
	if request.Price == nil {
		return apperror.ErrPriceRequired
	}
	return nil
}
