package router

import (
	"fmt"

	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/handler"
	"github.com/becardine/gestock-api/internal/repository"
	"github.com/becardine/gestock-api/internal/service"
	"github.com/go-chi/chi/v5"

	_ "github.com/becardine/gestock-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes(router *chi.Mux) {
	// initializa handler
	err := handler.Initialize()
	if err != nil {
		panic(err)
	}

	basePath := "/api/v1"

	usersPath := fmt.Sprintf("%s/users", basePath)

	db := config.GetDB()

	// repository
	userRepository := repository.NewUserRepository(db)

	// service
	userService := service.NewUserService(userRepository)

	// handler
	handler := handler.NewUserHandler(userService)

	// user routes
	router.Route(usersPath, func(r chi.Router) {
		r.Post("/", handler.CreateUser)
	})

	// Swagger routes
	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

}
