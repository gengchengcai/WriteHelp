package model

import "gorm.io/gorm"

type ChapterSetting struct {
	gorm.Model
	//章节内容设定
	ChapterContext string `json:"chaptercontext" form:"chaptercontext" gorm:"type:text; "`
	//章节简介
	ChapterIntroduction string `json:"chapterintroduction" form:"chapterintroduction" gorm:"type:tinytext; "`
	ChapterID           uint   `json:"chapterid" form:"chapterid"`
}
