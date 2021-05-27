package model

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	BookID          uint             `json:"bookid" form:"bookid"`
	ChapterSettings []ChapterSetting `json:"chaptersettings" form:"settings" `

	CapterName string `json:"captername" form:"captername" gorm:"type:varchar(64);not null;unique"`
}
