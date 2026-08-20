package product

import (
	"context"
	"testing"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

type fakeProductRepository struct {
	createRequest repositorymodel.CreateProductRequestModel
	createID      string
	createErr     error
	updateID      string
	updateErr     error
}

func (repository *fakeProductRepository) CreateProduct(executionContext context.Context, request *repositorymodel.CreateProductRequestModel, response *repositorymodel.CreateProductModel) error {
	if request != nil {
		repository.createRequest = *request
	}
	if repository.createErr != nil {
		return repository.createErr
	}
	response.ID = repository.createID
	response.Name = request.Name
	return nil
}

func (repository *fakeProductRepository) UpdateProduct(executionContext context.Context, request *repositorymodel.UpdateProductRequestModel) error {
	repository.updateID = request.ID
	return repository.updateErr
}

func TestCreateProductService_Create(t *testing.T) {
	price := 99.0
	service := NewProductService(&fakeProductRepository{createID: "p-1"})

	var response servicemodel.CreateProductResponseModel
	err := service.Create(context.Background(), &servicemodel.CreateProductRequestModel{
		Name:  "Green tea",
		Price: &price,
	}, &response)
	if err != nil {
		t.Fatal(err)
	}
	if response.ID != "p-1" || response.Name != "Green tea" {
		t.Fatalf("got %+v", response)
	}
}

func TestCreateProductService_RepositoryNotConfigured(t *testing.T) {
	service := NewProductService(nil)
	err := service.Create(context.Background(), &servicemodel.CreateProductRequestModel{}, &servicemodel.CreateProductResponseModel{})
	if err != apperror.ErrCreateProductRepositoryNotConfigured {
		t.Fatalf("got %v", err)
	}
}
