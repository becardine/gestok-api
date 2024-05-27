package main

import (
	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/router"
)

var (
	logger *config.Logger
)

// @title Gestok API
// @version 1.0
// @description This is a simple API for managing stock
// @termsOfService http://swagger.io/terms/

// @contact becardine
// @contact.email becardiine@gmail.com

// @license MIT

// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// initialize router
	router.InitializeRouter()
}
