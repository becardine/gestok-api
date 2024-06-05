package router

import (
	"fmt"
	"net/http"

	"github.com/becardine/gestock-api/config"
	"github.com/go-chi/chi/v5"
)

func InitializeRouter() {
	r := chi.NewRouter()

	// initialize routes api
	initializeRoutes(r)

	webServerPort := config.GetEnv().WebServerPort
	if webServerPort == "" {
		webServerPort = "8080"
	}

	config.GetLogger("router").Infof("web server running on port %s", webServerPort)

	addr := fmt.Sprintf("0.0.0.0:%s", webServerPort)
	http.ListenAndServe(addr, r)
}
