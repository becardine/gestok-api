package wire

import (
	"database/sql"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/application/handler"
	"github.com/becardine/gestock-api/internal/domain/repository"
	"github.com/becardine/gestock-api/internal/domain/service"
	infra "github.com/becardine/gestock-api/internal/infra/repository"
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

func InitializeCouponHandler() (*handler.CouponHandler, error) {
	db := DBProvider()

	var couponRepository repository.CouponRepositoryInterface
	couponRepository = infra.NewCouponRepository(db)
	couponService := service.NewCouponService(couponRepository)
	couponHandler := handler.NewCouponHandler(couponService)

	return couponHandler, nil
}

func InitializeConfig() error {
	return config.Init()
}

func DBProvider() *sql.DB {
	return config.GetDB()
}
