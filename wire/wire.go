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

func InitializeBrandHandler() (*handler.BrandHandler, error) {
	db := DBProvider()

	var brandRepository repository.BrandRepositoryInterface
	brandRepository = infra.NewBrandRepository(db)
	brandService := service.NewBrandService(brandRepository)
	brandHandler := handler.NewBrandHandler(brandService)

	return brandHandler, nil
}

func InitializeCategoryHandler() (*handler.CategoryHandler, error) {
	db := DBProvider()

	var categoryRepository repository.CategoryRepositoryInterface
	categoryRepository = infra.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	return categoryHandler, nil
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
