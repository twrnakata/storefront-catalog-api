package main

import (
	"log"

	httproute "github.com/twrnakata/storefront-catalog-api/internal/http/route"
	repositoryproduct "github.com/twrnakata/storefront-catalog-api/internal/repository/product"
	serviceproduct "github.com/twrnakata/storefront-catalog-api/internal/service/product"
	"github.com/twrnakata/storefront-catalog-api/pkg/configuration"
	"github.com/twrnakata/storefront-catalog-api/pkg/postgres"
)

func main() {
	if err := configuration.InitConfig(); err != nil {
		log.Fatal(err)
	}

	database, err := postgres.Connect(configuration.Env.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	createProductRepository, err := repositoryproduct.NewProductCreateRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	createProductService := &serviceproduct.CreateProductService{
		Repository: createProductRepository,
	}

	application := httproute.NewApp(createProductService)
	log.Printf("listening on :%s", configuration.Env.PORT)
	log.Fatal(application.Listen(":" + configuration.Env.PORT))
}
