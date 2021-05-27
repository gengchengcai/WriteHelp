package model

import "gorm.io/gorm"

//这个是每本书的章节结构体，包含该章节的简介，所属于哪本书，章节设定，可以有多个设定和章节的名字
type Chapter struct {
	gorm.Model
	BookID uint `json:"bookid" form:"bookid"`

	ChapterContext string `json:"chaptercontext" form:"chaptercontext" gorm:"type:text; "`
	//章节简介
	ChapterIntroduction string `json:"chapterintroduction" form:"chapterintroduction" gorm:"type:tinytext; "`
	CapterName          string `json:"captername" form:"captername" gorm:"type:varchar(64);not null;unique"`
}
