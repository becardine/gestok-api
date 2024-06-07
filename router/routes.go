package router

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	"github.com/becardine/gestock-api/config"
	_ "github.com/becardine/gestock-api/docs"
	"github.com/becardine/gestock-api/wire"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes(router *chi.Mux) {
	// initializa handler
	// err := handler.Initialize()
	// if err != nil {
	// 	panic(err)
	// }

	basePath := "/api/v1"

	usersPath := fmt.Sprintf("%s/users", basePath)

	db := config.GetDB()

	// repository
	// userRepository := repository.NewUserRepository(db)

	// service
	// userService := service.NewUserService(userRepository)

	// handler
	// handler := handler.NewUserHandler(userService)

	// user routes
	router.Route(usersPath, func(r chi.Router) {
		// r.Post("/", handler.CreateUser)
		// r.Get("/{email}", handler.GetUserByEmail)
	})

	// product routes
	productHandler := wire.InitializeProductHandler(db)
	router.Route("/products", func(r chi.Router) {
		r.Mount("/", productHandler.Routes())
	})

	// Swagger routes
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

}
