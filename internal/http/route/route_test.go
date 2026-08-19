package route

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
)

type fakeCreateProductService struct{}

func (service *fakeCreateProductService) Create(executionContext context.Context, request *servicemodel.CreateProductRequestModel, response *servicemodel.CreateProductResponseModel) error {
	response.ID = "id-1"
	response.Name = request.Name
	return nil
}

type fakeUpdateProductService struct{}

func (service *fakeUpdateProductService) Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error {
	return nil
}

func TestPostProduct_Component(t *testing.T) {
	application := NewApp(&fakeCreateProductService{}, &fakeUpdateProductService{})
	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"Tea","price":10}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("status %d", response.StatusCode)
	}
	raw, _ := io.ReadAll(response.Body)
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatal(err)
	}
	if payload["successful"] != true {
		t.Fatalf("got %s", raw)
	}
}

func TestAPIDocs_Index(t *testing.T) {
	application := NewApp(&fakeCreateProductService{}, &fakeUpdateProductService{})
	request := httptest.NewRequest("GET", "/api-docs", nil)
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("status %d", response.StatusCode)
	}
	raw, _ := io.ReadAll(response.Body)
	if !bytes.Contains(raw, []byte("swagger-ui")) {
		t.Fatalf("missing swagger ui: %s", raw)
	}
	if !bytes.Contains(raw, []byte("English (EN)")) {
		t.Fatalf("missing language urls: %s", raw)
	}
}

func TestAPIDocs_OpenAPIFiles(t *testing.T) {
	application := NewApp(&fakeCreateProductService{}, &fakeUpdateProductService{})
	for _, path := range []string{"/api-docs/openapi.en.yaml", "/api-docs/openapi.th.yaml"} {
		request := httptest.NewRequest("GET", path, nil)
		response, err := application.Test(request, -1)
		if err != nil {
			t.Fatal(err)
		}
		if response.StatusCode != fiber.StatusOK {
			t.Fatalf("%s status %d", path, response.StatusCode)
		}
		raw, _ := io.ReadAll(response.Body)
		if !bytes.Contains(raw, []byte("openapi: 3.0.3")) {
			t.Fatalf("%s missing openapi version: %s", path, raw)
		}
	}
}

func TestUpdateProduct_Component(t *testing.T) {
	application := NewApp(&fakeCreateProductService{}, &fakeUpdateProductService{})
	productID := "11111111-1111-1111-1111-111111111111"
	request := httptest.NewRequest("PATCH", "/product/"+productID, bytes.NewBufferString(`{"description":null}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("status %d", response.StatusCode)
	}
}
