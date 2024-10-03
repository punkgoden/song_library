package http_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/punkgoden/song_library/app/internal/config"
	"github.com/punkgoden/song_library/app/internal/service"
	"github.com/punkgoden/song_library/app/pkg/logging"
)

type Controller struct {
	log logging.Logger
	sc  service.Service
	cfg config.Config
}

func NewController(
	log logging.Logger,
	sc service.Service,
	cfg config.Config,
) *Controller {
	return &Controller{
		log: log,
		sc:  sc,
		cfg: cfg,
	}
}

func (c *Controller) InitRoutes(r *gin.Engine) *gin.Engine {

	api := r.Group("api/")
	{
		r.Group("library/")
		{
			api.GET("song", c.GetSong)
			api.POST("song", c.CreateSong)
			api.PATCH("song", c.UpdateSong)
			api.DELETE("song", c.DeleteSong)
		}
		api.GET("healthcheck", c.Healthcheck)
	}

	return r
}
