package main

import (
	"github.com/gin-gonic/gin"
	"writer/controller"
	"writer/logger"
	"writer/tool"
)

func main() {

	app := gin.Default()
	tool.InitDB()
	logger.InitLogrus("G:\\goland\\writer\\Log\\serve.log")
	v1Group := app.Group("v1")
	Router(v1Group)
	app.Run()
}

func Router(engine *gin.RouterGroup) {
	new(controller.Routers).CreateRouters(engine)
}
