package config

import (
	"fmt"

	"github.com/becardine/gestock-api/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("PostgreSQL")

	// create connection
	env := GetEnv()

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=America/Sao_Paulo",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect database: %v", err)
		return nil, err
	}

	logger.Info("database connection established")

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.Errorf("failed to migrate database: %v", err)
		return nil, err
	}

	logger.Info("database migrated")

	return db, nil
}
