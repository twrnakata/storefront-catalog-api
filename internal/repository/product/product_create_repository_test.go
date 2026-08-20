package product

import (
	"context"
	"os"
	"testing"

	repositorymodel "github.com/twrnakata/storefront-catalog-api/internal/repository/product/model"
	"github.com/twrnakata/storefront-catalog-api/pkg/postgres"
)

func TestProductCreateRepository_CreateProduct(t *testing.T) {
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

	price := 12.5
	created := &repositorymodel.CreateProductModel{}
	err = repository.CreateProduct(context.Background(), &repositorymodel.CreateProductRequestModel{
		Name:  "Repo tea",
		Price: price,
	}, created)
	if err != nil {
		t.Fatal(err)
	}
	if created.ID == "" {
		t.Fatal("missing id")
	}
}
