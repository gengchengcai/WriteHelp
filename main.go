package main

import (
	"github.com/gin-gonic/gin"
	"writer/tool"
)

func main() {

	app := gin.Default()
	tool.InitDB()

	app.Run()
}
