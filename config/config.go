package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
	cfg    *conf
)

func Init() error {
	var err error

	// initialize environments
	cfg, err = LoadConfig(".")
	if err != nil {
		return fmt.Errorf("failed to load enviroments: %v", err)
	}

	// initialize db
	db, err = InitializePostgreSQL()
	if err != nil {
		return fmt.Errorf("failed to initialize db: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetEnv() *conf {
	return cfg
}

func GetDB() *gorm.DB {
	return db
}
