package database

import (
	"fmt"
	"github.com/punkgoden/song_library/app/internal/config"
	"github.com/punkgoden/song_library/app/internal/entity"
	"github.com/punkgoden/song_library/app/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config, log *logging.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.AppConfig.PostgresSQL.Username,
		cfg.AppConfig.PostgresSQL.Password,
		cfg.AppConfig.PostgresSQL.Host,
		cfg.AppConfig.PostgresSQL.Port,
		cfg.AppConfig.PostgresSQL.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
		return nil, err
	}

	if err := migrate(db); err != nil {
		log.Errorln("Error migrate database")
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.Song{}); err != nil {
		return err
	}
	return nil
}
