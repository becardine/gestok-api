package main

import "github.com/becardine/gestock-api/config"

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
}
