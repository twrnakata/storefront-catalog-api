package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	apidocs "github.com/twrnakata/storefront-catalog-api/api-docs"
	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	"github.com/twrnakata/storefront-catalog-api/internal/http/handler"
)

func NewApp(productService domainproduct.ProductService) *fiber.App {
	application := fiber.New()
	productHandler := handler.NewProductHandler(productService)
	application.Post("/product", productHandler.Create)
	application.Patch("/product/:id", productHandler.Update)
	application.Use("/api-docs", filesystem.New(filesystem.Config{
		Root:   http.FS(apidocs.Files),
		Index:  "index.html",
		MaxAge: 0,
	}))
	return application
}
