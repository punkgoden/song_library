package service

import (
	"github.com/punkgoden/song_library/app/internal/storage"
	"github.com/punkgoden/song_library/app/pkg/config"
	"github.com/punkgoden/song_library/app/pkg/logging"
)

type Service struct {
	st  storage.Storage
	log logging.Logger
	cfg config.Config
}

func NewService(
	log logging.Logger,
	st storage.Storage,
	cfg config.Config,
) *Service {
	return &Service{
		st:  st,
		log: log,
		cfg: cfg,
	}
}
