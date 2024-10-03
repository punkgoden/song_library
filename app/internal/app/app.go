package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/punkgoden/song_library/app/internal/config"
	http_controller "github.com/punkgoden/song_library/app/internal/controller/http"
	database "github.com/punkgoden/song_library/app/internal/db/postgres"
	sc "github.com/punkgoden/song_library/app/internal/service"
	st "github.com/punkgoden/song_library/app/internal/storage"
	"github.com/punkgoden/song_library/app/pkg/logging"
	"github.com/punkgoden/song_library/app/pkg/server"
)

func RunApplication(saveToFile bool) {
	// Init Logger
	logging.Init(false)
	log := logging.GetLogger()
	log.Infoln("Connect logger successfully!")

	// Init Config
	cfg := config.GetConfig()
	log.Infoln("Connect config successfully!")

	ctx := context.Background()

	db, err := database.NewPostgresDB(cfg, &log)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Infoln("Connect database successfully!")

	gin.SetMode(cfg.AppConfig.GinMode)

	storage := st.NewStorage(log, db)
	log.Infoln("Connect storage successfully!")
	service := sc.NewService(log, *storage, *cfg)
	log.Infoln("Connect service successfully!")

	ginRouter := gin.New()
	httpController := http_controller.NewController(ctx, log, *service, *cfg)
	handlers := httpController.InitRoutes(ginRouter)
	log.Infoln("Connect GIN successfully!")

	server.RunServer(log, handlers, cfg.AppConfig.HttpAddr)
}
