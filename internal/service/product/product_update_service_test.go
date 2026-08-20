package product

import (
	"context"
	"testing"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/optional"
)

func TestUpdateProductService_Update(t *testing.T) {
	repository := &fakeProductRepository{}
	service := NewProductService(repository)
	err := service.Update(context.Background(), &servicemodel.UpdateProductRequestModel{
		ID:   "id-1",
		Name: optional.From("Tea"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if repository.updateID != "id-1" {
		t.Fatalf("got %s", repository.updateID)
	}
}

func TestUpdateProductService_NotFound(t *testing.T) {
	service := NewProductService(&fakeProductRepository{updateErr: domainproduct.ErrProductNotFound})
	err := service.Update(context.Background(), &servicemodel.UpdateProductRequestModel{ID: "missing"})
	if err != domainproduct.ErrProductNotFound {
		t.Fatalf("got %v", err)
	}
}
