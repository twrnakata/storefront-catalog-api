package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	servicemodel "github.com/twrnakata/storefront-catalog-api/internal/service/product/model"
)

type fakeUpdateProductService struct {
	err error
	id  string
}

func (service *fakeUpdateProductService) Update(executionContext context.Context, request *servicemodel.UpdateProductRequestModel) error {
	service.id = request.ID
	return service.err
}

func TestProductUpdateHandler_Update(t *testing.T) {
	handler := &ProductUpdateHandler{UpdateService: &fakeUpdateProductService{}}
	application := fiber.New()
	application.Patch("/product/:id", handler.Update)

	productID := "11111111-1111-1111-1111-111111111111"
	request := httptest.NewRequest("PATCH", "/product/"+productID, bytes.NewBufferString(`{"name":"Tea"}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("status %d", response.StatusCode)
	}
}

func TestProductUpdateHandler_InvalidID(t *testing.T) {
	handler := &ProductUpdateHandler{UpdateService: &fakeUpdateProductService{}}
	application := fiber.New()
	application.Patch("/product/:id", handler.Update)

	request := httptest.NewRequest("PATCH", "/product/not-a-uuid", bytes.NewBufferString(`{"name":"Tea"}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("status %d", response.StatusCode)
	}
}

func TestProductUpdateHandler_NoFields(t *testing.T) {
	handler := &ProductUpdateHandler{UpdateService: &fakeUpdateProductService{}}
	application := fiber.New()
	application.Patch("/product/:id", handler.Update)

	productID := "11111111-1111-1111-1111-111111111111"
	request := httptest.NewRequest("PATCH", "/product/"+productID, bytes.NewBufferString(`{}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("status %d", response.StatusCode)
	}
	raw, _ := io.ReadAll(response.Body)
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatal(err)
	}
	if payload["error_code"] != "at least one field is required" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductUpdateHandler_NotFound(t *testing.T) {
	handler := &ProductUpdateHandler{UpdateService: &fakeUpdateProductService{err: domainproduct.ErrProductNotFound}}
	application := fiber.New()
	application.Patch("/product/:id", handler.Update)

	productID := "11111111-1111-1111-1111-111111111111"
	request := httptest.NewRequest("PATCH", "/product/"+productID, bytes.NewBufferString(`{"name":"Tea"}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusNotFound {
		t.Fatalf("status %d", response.StatusCode)
	}
}
