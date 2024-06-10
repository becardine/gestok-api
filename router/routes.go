package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	"github.com/becardine/gestock-api/config"
	_ "github.com/becardine/gestock-api/docs"
	"github.com/becardine/gestock-api/internal/application/handler"
	"github.com/becardine/gestock-api/wire"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes(router *chi.Mux) {
	config.GetLogger("router").Info("initializing routes")

	basePath := "/api/v1"

	usersPath := fmt.Sprintf("%s/users", basePath)

	// user routes
	router.Route(usersPath, func(r chi.Router) {
		// r.Post("/", handler.CreateUser)
		// r.Get("/{email}", handler.GetUserByEmail)
	})

	router.Route(basePath, func(r chi.Router) {
		// product routes
		productHandler, err := wire.InitializeProductHandler()
		if err != nil {
			config.GetLogger("router").Errorf("error while initializing product handler: %v", err)
		}
		r.Route("/products", productHandler.Routes)

		// test routes
		testHandler := handler.NewTestHandler()
		r.Route("/test", testHandler.Routes)
	})

	// Swagger routes
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

}
