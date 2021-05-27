package model

import "gorm.io/gorm"

type PlotSetting struct {
	gorm.Model
	BookID              uint   `json:"bookid" form:"bookid"`
	PlotSettingContexts string `json:"plotsettingcontexts" form:"plotsettingcontexts" gorm:"type:text"`
}
