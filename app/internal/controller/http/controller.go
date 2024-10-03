package http_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/punkgoden/song_library/app/internal/config"
	"github.com/punkgoden/song_library/app/internal/service"
	"github.com/punkgoden/song_library/app/pkg/logging"
	"github.com/punkgoden/song_library/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title      		Song Library
// @version     	0.1
// @description 	Test project

// @contact.name   	go_rshok
// @contact.url    	https://t.me/go_rshok

// @BasePath  		/api/

type Controller struct {
	ctx context.Context
	log logging.Logger
	sc  service.Service
	cfg config.Config
}

func NewController(
	ctx context.Context,
	log logging.Logger,
	sc service.Service,
	cfg config.Config,
) *Controller {
	return &Controller{
		ctx: ctx,
		log: log,
		sc:  sc,
		cfg: cfg,
	}
}

func (c *Controller) InitRoutes(r *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api"

	doc := ginSwagger.URL("http://localhost:8000/api/swagger/doc.json")
	api := r.Group("api/")
	{
		api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), doc))

		api.GET("songs", c.GetSongs)
		api.GET("song", c.GetSong)
		api.GET("text-song", c.GetTextSong)
		api.POST("song", c.CreateSong)
		api.PATCH("song", c.UpdateSong)
		api.DELETE("song", c.DeleteSong)

		api.GET("healthcheck", c.Healthcheck)
	}

	return r
}
