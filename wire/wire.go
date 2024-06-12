package wire

import (
	"database/sql"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/application/handler"
	"github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/domain/service"
	infra "github.com/becardine/gestock-api/internal/infra/repository"
	"github.com/becardine/gestock-api/internal/mocks"
)

func InitializeProductHandler() (*handler.ProductHandler, error) {
	db := DBProvider()

	var productRepository repository.ProductRepository
	productRepository = infra.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	return productHandler, nil
}

func InitializeConfig() error {
	return config.Init()
}

func DBProvider() *sql.DB {
	return config.GetDB()
}

func MockProductService() service.ProductService {
	return new(mocks.ProductServiceMock)
}
