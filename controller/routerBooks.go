package controller

import (
	"github.com/gin-gonic/gin"
	"writer/dao"
	"writer/logger"
	"writer/model"
	"writer/tool"
)

func (r *Routers) GetBooksInfo(c *gin.Context) {
	var user model.Author
	db := tool.GetOrmDB()
	err := c.ShouldBind(&user)
	if err != nil {
		logger.LogError(map[string]interface{}{"msg": "获取数据失败"}).Info("获取图书信息")
		return
	}
	info, ok := dao.IsAuthorDao(db, user)
	if !ok {
		logger.LogError(map[string]interface{}{"msg": "没有该作者"})
		c.JSON(200, gin.H{
			"code":    "40000",
			"data":    "",
			"message": "该作者还没有注册",
		})
		return
	}
	BooksNum := len(info.Books)
	NewUser := model.Author{
		UserName: info.UserName,
		Books:    info.Books,
	}
	logger.LogInfo(map[string]interface{}{"msg": "getbooksinfo success!"})
	c.JSON(200, gin.H{
		"code":    "200",
		"data":    NewUser,
		"array":   BooksNum,
		"message": "获取成功",
	})
}
