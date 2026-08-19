package route

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	apidocs "github.com/twrnakata/storefront-catalog-api/api-docs"
	domainproduct "github.com/twrnakata/storefront-catalog-api/internal/domain/product"
	"github.com/twrnakata/storefront-catalog-api/internal/http/handler"
)

func NewApp(createProductService domainproduct.CreateProductService, updateProductService domainproduct.UpdateProductService) *fiber.App {
	application := fiber.New()
	productCreateHandler := &handler.ProductCreateHandler{CreateService: createProductService}
	productUpdateHandler := &handler.ProductUpdateHandler{UpdateService: updateProductService}
	application.Post("/product", productCreateHandler.Create)
	application.Patch("/product/:id", productUpdateHandler.Update)
	application.Use("/api-docs", filesystem.New(filesystem.Config{
		Root:   http.FS(apidocs.Files),
		Index:  "index.html",
		MaxAge: 0,
	}))
	return application
}
