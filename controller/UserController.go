package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (userController *UserController) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello gin",
	})
}
