package controller

import (
	"github.com/gin-gonic/gin"
	"writer/logger"
	"writer/model"
)

func (r *Routers) CreateRouters(engine *gin.RouterGroup) {
	engine.POST("users", r.Register)
	engine.GET("users", r.Login)
}

func (r *Routers) Register(c *gin.Context) {
	var user model.Author
	if err := c.ShouldBind(&user); err != nil {

	}
}

func (r *Routers) Login(c *gin.Context) {

	logger.LogInfo(map[string]interface{}{"test": 1}).Info("this is a test")
	c.JSON(200, gin.H{
		"test": "test",
	})
}
