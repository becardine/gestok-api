package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitializePostgreSQL() (*sql.DB, error) {
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
	db, err := sql.Open("postgres", dns)
	if err != nil {
		logger.Errorf("failed to connect database: %v", err)
		return nil, err
	}

	logger.Info("database connection established")

	return db, nil
}
