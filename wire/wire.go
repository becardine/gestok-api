package wire

import (
	"database/sql"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/application/handler"
	"github.com/becardine/gestock-api/internal/domain/service"
	"github.com/becardine/gestock-api/internal/infra/repository"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	repository.NewProductRepository,
)

func InitializeProductHandler() (*handler.ProductHandler, error) {
	wire.Build(
		DBProvider,
		setRepositoryDependency,
		service.NewProductService,
		handler.NewProductHandler,
	)
	return &handler.ProductHandler{}, nil
}

func DBProvider() (*sql.DB, error) {
	return config.InitializePostgreSQL()
}
