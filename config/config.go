package config

import (
	"database/sql"
	"fmt"
)

var (
	db     *sql.DB
	logger *Logger
	cfg    *Conf
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

func GetEnv() *Conf {
	return cfg
}

func GetDB() *sql.DB {
	return db
}
