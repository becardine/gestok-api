package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeMySQL() (*sql.DB, error) {
	logger := GetLogger("MySQL")

	// create connection
	env := GetEnv()

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DBUser,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)
	logger.Infof("connecting to database: %s", dns)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		logger.Errorf("failed to connect database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Errorf("failed to ping database: %v", err)
		return nil, err
	}

	logger.Info("database connection established")

	return db, nil
}
