package handler

import (
	"github.com/becardine/gestock-api/config"
)

var (
	logger *config.Logger
)

func Initialize() error {
	logger = config.GetLogger("handler")
	return nil
}
