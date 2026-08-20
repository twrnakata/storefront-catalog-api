package product

import (
	"context"
	"os"
	"testing"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/optional"
	"github.com/twrnakata/storefront-catalog-api/pkg/postgres"
)

func TestProductUpdateRepository_UpdateProduct(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("DATABASE_URL not set")
	}

	database, err := postgres.Connect(databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	repository, err := NewProductRepository(database)
	if err != nil {
		t.Fatal(err)
	}

	description := "before update"
	salePrice := 8.5
	created := &repositorymodel.CreateProductModel{}
	err = repository.CreateProduct(context.Background(), &repositorymodel.CreateProductRequestModel{
		Name:        "Repo tea",
		Description: &description,
		SalePrice:   &salePrice,
		Price:       12.5,
	}, created)
	if err != nil {
		t.Fatal(err)
	}

	updatedName := "Updated repo tea"
	updatedPrice := 15.0
	err = repository.UpdateProduct(context.Background(), &repositorymodel.UpdateProductRequestModel{
		ID:    created.ID,
		Name:  optional.From(updatedName),
		Price: optional.From(updatedPrice),
	})
	if err != nil {
		t.Fatal(err)
	}

	var record ProductRecord
	err = database.WithContext(context.Background()).Where("id = ?", created.ID).First(&record).Error
	if err != nil {
		t.Fatal(err)
	}
	if record.Name != updatedName {
		t.Fatalf("expected name %q, got %q", updatedName, record.Name)
	}
	if record.Price != updatedPrice {
		t.Fatalf("expected price %v, got %v", updatedPrice, record.Price)
	}
	if record.Description == nil || *record.Description != description {
		t.Fatalf("expected description %q to remain unchanged, got %#v", description, record.Description)
	}
	if record.SalePrice == nil || *record.SalePrice != salePrice {
		t.Fatalf("expected sale price %v to remain unchanged, got %#v", salePrice, record.SalePrice)
	}
}

func TestProductUpdateRepository_ProductNotFound(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("DATABASE_URL not set")
	}

	database, err := postgres.Connect(databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	repository, err := NewProductRepository(database)
	if err != nil {
		t.Fatal(err)
	}

	err = repository.UpdateProduct(context.Background(), &repositorymodel.UpdateProductRequestModel{
		ID:   "missing-id",
		Name: optional.From("Tea"),
	})
	if err != domainproduct.ErrProductNotFound {
		t.Fatalf("expected %v, got %v", domainproduct.ErrProductNotFound, err)
	}
}
