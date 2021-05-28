package model

import "gorm.io/gorm"

// 用作注册和登录的表：用户，密码，创建书就在里面加入。
type Author struct {
	gorm.Model
	UserName string `json:"username" form:"username" gorm:"type:varchar(64); not null; unique" binding:"required" `
	PassWord string `json:"password" form:"password" gorm:"type:varchar(64); not null" binding: "required" `
	Books    []Book `json:"books" form:"books" `
}

func (Author) TableName() string {
	return "author"
}
