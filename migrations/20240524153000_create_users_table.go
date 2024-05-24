package migrations

import (
	"github.com/becardine/gestock-api/config"
	"github.com/becardine/gestock-api/internal/entity"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	logger := config.GetLogger("migrations")

	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.Errorf("failed to migrate database: %v", err)
		return err
	}
	return nil
}

func Down(db *gorm.DB) error {
	logger := config.GetLogger("migrations")

	err := db.Migrator().DropTable(&entity.User{})
	if err != nil {
		logger.Errorf("failed to drop table: %v", err)
		return err
	}
	return nil
}
