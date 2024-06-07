package wire

import (
	"database/sql"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/application/handler"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/infra/repository"
	"github.com/google/wire"
)

func InitializeProductHandler(db *sql.DB) *handler.ProductHandler {
	wire.Build(
		repository.NewProductRepository,
		service.NewProductService,
		handler.NewProductHandler,
	)
	return &handler.ProductHandler{}
}

func DBProvider() (*sql.DB, error) {
	return config.InitializePostgreSQL()
}
