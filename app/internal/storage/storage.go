package storage

import (
	"github.com/punkgoden/song_library/app/pkg/logging"

	"gorm.io/gorm"
)

type Storage struct {
	db  *gorm.DB
	log logging.Logger
}

func NewStorage(log logging.Logger, db *gorm.DB) *Storage {
	return &Storage{
		db:  db,
		log: log,
	}
}
