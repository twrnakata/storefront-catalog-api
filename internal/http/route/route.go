package route

import (
	"github.com/gofiber/fiber/v2"

	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	"github.com/twrnakata/storefront-catalog-api/internal/http/handler"
)

func NewApp(createProductService domainproduct.CreateProductService) *fiber.App {
	application := fiber.New()
	productCreateHandler := &handler.ProductCreateHandler{CreateService: createProductService}
	application.Post("/product", productCreateHandler.Create)
	return application
}
