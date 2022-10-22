package main

import (
	_ "gin-demo/config"
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	routers.RegisterRouter(engine)

	engine.Run(":9090")
}
