package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"writer/dao"
	"writer/logger"
	"writer/model"
	"writer/tool"
)

func (r *Routers) CreateRouters(engine *gin.RouterGroup) {
	engine.POST("register", r.Register)
	engine.POST("login", r.Login)
}

func (r *Routers) Register(c *gin.Context) {
	logger.LogInfo(map[string]interface{}{}).Info("用户进行注册")
	var user model.Author
	if err := c.ShouldBind(&user); err != nil {
		logger.LogError(map[string]interface{}{}).Info("注册出现问题")
		log.Fatal(err.Error())
		return
	}
	db := tool.GetOrmDB()
	if dao.UserRegisterDao(db, user) {
		logger.LogInfo(map[string]interface{}{}).Info("注册成功")
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "注册成功",
			"data":    user,
		})
	} else {
		logger.LogInfo(map[string]interface{}{}).Info("创建数据出现问题，账号重复可能性极大")
		c.JSON(200, gin.H{
			"code":    "422",
			"message": "账号重复",
		})
	}

}

func (r *Routers) Login(c *gin.Context) {
	logger.LogInfo(map[string]interface{}{}).Info("用户进行登录")
	var user model.Author
	db := tool.GetOrmDB()
	err := c.ShouldBind(&user)
	if err != nil {
		logger.LogError(map[string]interface{}{}).Info("用户登录时出错")
	}
	if !dao.UserLoginDao(db, user) {
		c.JSON(200, gin.H{
			"code":    "422",
			"message": "账号或者密码出现问题",
		})
		logger.LogInfo(map[string]interface{}{}).Info("尝试登录失败")
		return
	}
	logger.LogInfo(map[string]interface{}{}).Info("用户注册成功")
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "登录成功",
		"data":    "allow", //这块应该给一个token
	})
}
