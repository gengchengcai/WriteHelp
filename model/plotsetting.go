package model

import "gorm.io/gorm"

type PlotSetting struct {
	gorm.Model
	BookID uint `json:"bookid" form:"bookid"`
	//剧情设定
	PlotSettingContexts string `json:"plotsettingcontexts" form:"plotsettingcontexts" gorm:"type:text"`
}
