package http_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) Healthcheck(ctx *gin.Context) {
	var obj struct {
		Msg string `json:"msg"`
	}
	obj.Msg = "Hello Song Library"
	ctx.JSON(http.StatusOK, &obj)
}
