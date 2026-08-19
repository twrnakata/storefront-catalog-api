package product

import (
	"context"
	"testing"

	repositoryproduct "github.com/twrnakata/storefront-catalog-api/internal/repository/product"
	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/optional"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
)

type fakeUpdateProductRepository struct {
	err error
	id  string
}

func (repository *fakeUpdateProductRepository) UpdateProduct(executionContext context.Context, request *repositorymodel.UpdateProductRequestModel) error {
	repository.id = request.ID
	return repository.err
}

func TestUpdateProductService_Update(t *testing.T) {
	repository := &fakeUpdateProductRepository{}
	service := &UpdateProductService{Repository: repository}
	err := service.Update(context.Background(), &servicemodel.UpdateProductRequestModel{
		ID:   "id-1",
		Name: optional.From("Tea"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if repository.id != "id-1" {
		t.Fatalf("got %s", repository.id)
	}
}

func TestUpdateProductService_NotFound(t *testing.T) {
	service := &UpdateProductService{Repository: &fakeUpdateProductRepository{err: repositoryproduct.ErrProductNotFound}}
	err := service.Update(context.Background(), &servicemodel.UpdateProductRequestModel{ID: "missing"})
	if err != domainproduct.ErrProductNotFound {
		t.Fatalf("got %v", err)
	}
}
