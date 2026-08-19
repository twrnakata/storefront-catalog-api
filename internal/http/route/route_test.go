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

func TestPostProduct_Component(t *testing.T) {
	application := NewApp(&fakeCreateProductService{})
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
