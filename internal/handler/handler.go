package handler

import (
	"github.com/becardine/gestock-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize() error {
	logger = config.GetLogger("handler")
	var err error
	db, err = config.InitializePostgreSQL()
	if err != nil {
		return err
	}
	return nil
}
