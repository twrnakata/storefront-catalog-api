package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestProductCreateHandler_Create(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{createID: "abc"})
	application := fiber.New()
	application.Post("/product", handler.Create)

	body := `{"name":"Green tea","price":45.5,"description":null,"sale_price":null}`
	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(body))
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
	data := payload["data"].(map[string]any)
	if data["data1"] != "abc" || data["data2"] != "Green tea" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductCreateHandler_InvalidJSON_Returns400(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("status %d", response.StatusCode)
	}
}

func TestProductCreateHandler_NameRequired_Returns400(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"  ","price":10}`))
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
	if payload["error_code"] != "name is required" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductCreateHandler_PriceRequired_Returns400(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"Tea"}`))
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
	if payload["error_code"] != "price is required" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductCreateHandler_NegativePrice_Returns400(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"Tea","price":-1}`))
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
	if payload["error_code"] != "price must be greater than or equal to 0" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductCreateHandler_NegativeSalePrice_Returns400(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"Tea","price":10,"sale_price":-1}`))
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
	if payload["error_code"] != "sale_price must be greater than or equal to 0" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductCreateHandler_SalePriceExceedsPrice_Returns400(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"Tea","price":10,"sale_price":20}`))
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
	if payload["error_code"] != "sale_price must not exceed price" {
		t.Fatalf("got %s", raw)
	}
}

func TestProductCreateHandler_ZeroPrice_Returns200(t *testing.T) {
	handler := newTestProductHandler(&fakeProductService{createID: "abc"})
	application := fiber.New()
	application.Post("/product", handler.Create)

	request := httptest.NewRequest("POST", "/product", bytes.NewBufferString(`{"name":"Free sample","price":0}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := application.Test(request, -1)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("status %d", response.StatusCode)
	}
}
