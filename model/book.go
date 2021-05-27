package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	//BooKID uint `json:"bookid" form:"bookid" gorm:"type:int(10); unique; not null;"`
	BookName string `json:"bookname" form:"bookname" gorm:"type:varchar(64); not null; unique"`
	//关联了作者的主键
	AuthorID uint `json:"authorid" form:"authorid" gorm:"type:uint;not null;unique"`
	//书籍大纲
	OutlineSetting   string `json:"outlinesetting" form:"outlinesetting" gorm:"type:tinytext "`
	BookID           uint   `json:"bookid" form:"bookid"`
	BookIntroduction string `json:"bookintroduction" form:"bookintroduction" gorm:"type:tinytext"`
	Type             string `json:"type" form:"type" gorm:"type:varchar(12);not null"`
	//人物设定
	RoleSettings []Rolesetting `json:"rolesettings" form:"rolesettings" `
	//章节
	Chapters []Chapter `json:"chapters" form:"chapters"`
	//剧情设定
	PlotSettings []PlotSetting `json:"plotsetting" form:"plotsettings" `
	//灵感
	Inspirations []Inspiration `json:"inspirations" form:"inspirations" `
}
