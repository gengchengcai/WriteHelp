package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	UserName string `json:"username" form:"username" gorm:"type:varchar(64); not null; unique"`
	PassWord string `json:"password" form:"password" gorm:"type:varchar(64); not null; unique"`
	Books    []Book `json:"books" form:"books" `
}

func (Author) TableName() string {
	return "author"
}
