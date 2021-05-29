package dao

import (
	"gorm.io/gorm"
	"writer/logger"
	"writer/model"
)

func IsAuthorDao(db *gorm.DB, user model.Author) (model.Author, bool) {
	var NewUser model.Author
	db.Where("username = ?", user.UserName).Preload("Books").Preload("Chapters").First(&NewUser)
	if NewUser.ID == 0 {
		logger.LogInfo(map[string]interface{}{"msg": "该作者没有注册"}).Info("获取图书信息")
		return model.Author{}, false
	}
	return NewUser, true
}
