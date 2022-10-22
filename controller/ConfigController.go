package controller

import (
	"gin-demo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConfigController struct {
}

func (c *ConfigController) GetConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, config.GlobalConfig)
}
