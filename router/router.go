package router

import (
	"fmt"
	"net/http"

	"github.com/becardine/gestock-api/config"
	"github.com/go-chi/chi/v5"
)

func Init() {
	r := chi.NewRouter()
	initializeRoutes(r)

	webServerPort := config.GetEnv().WebServerPort
	if webServerPort == "" {
		webServerPort = "8080"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", webServerPort)
	http.ListenAndServe(addr, r)
}
