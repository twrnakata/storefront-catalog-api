package handler

import (
	"context"

	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
)

type fakeProductService struct {
	createID   string
	createErr  error
	updateErr  error
}

func (service *fakeProductService) Create(executionContext context.Context, request *servicemodel.CreateProductRequestModel, response *servicemodel.CreateProductResponseModel) error {
	if service.createErr != nil {
		return service.createErr
	}
	response.ID = service.createID
	response.Name = request.Name
	return nil
}

func (service *fakeProductService) Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error {
	return service.updateErr
}

func newTestProductHandler(service *fakeProductService) *ProductHandler {
	return NewProductHandler(service)
}
