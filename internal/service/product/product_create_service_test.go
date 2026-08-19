package product

import (
	"context"
	"testing"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

type fakeCreateProductRepository struct {
	request repositorymodel.CreateProductRequestModel
	id      string
	err     error
}

func (repository *fakeCreateProductRepository) CreateProduct(executionContext context.Context, request *repositorymodel.CreateProductRequestModel, response *repositorymodel.CreateProductModel) error {
	if request != nil {
		repository.request = *request
	}
	if repository.err != nil {
		return repository.err
	}
	response.ID = repository.id
	response.Name = request.Name
	return nil
}

func TestCreateProductService_Create(t *testing.T) {
	price := 99.0
	repository := &fakeCreateProductRepository{id: "p-1"}
	service := &CreateProductService{Repository: repository}

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
	service := &CreateProductService{}
	err := service.Create(context.Background(), &servicemodel.CreateProductRequestModel{}, &servicemodel.CreateProductResponseModel{})
	if err != apperror.ErrCreateProductRepositoryNotConfigured {
		t.Fatalf("got %v", err)
	}
}
