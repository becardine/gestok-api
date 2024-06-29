package router

import (
	"fmt"

	"github.com/becardine/gestock-api/wire"

	"github.com/go-chi/chi/v5"

	"github.com/becardine/gestock-api/config"
	_ "github.com/becardine/gestock-api/docs"
	"github.com/becardine/gestock-api/internal/application/handler"
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
		r.Route("/products", productHandler.RegisterRoutes) // /api/v1/products

		// brand routes
		brandHandler, err := wire.InitializeBrandHandler()
		if err != nil {
			config.GetLogger("router").Errorf("error while initializing brand handler: %v", err)
		}
		r.Route("/brands", brandHandler.RegisterRoutes) // /api/v1/brands

		// category routes
		categoryHandler, err := wire.InitializeCategoryHandler()
		if err != nil {
			config.GetLogger("router").Errorf("error while initializing category handler: %v", err)
		}
		r.Route("/categories", categoryHandler.RegisterRoutes) // /api/v1/categories

		// coupon routes
		couponHandler, err := wire.InitializeCouponHandler()
		if err != nil {
			config.GetLogger("router").Errorf("error while initializing coupon handler: %v", err)
		}
		r.Route("/coupons", couponHandler.RegisterRoutes) // /api/v1/coupons

		// test routes
		testHandler := handler.NewTestHandler()
		r.Route("/test", testHandler.Routes)
	})

	// Swagger routes
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

}
