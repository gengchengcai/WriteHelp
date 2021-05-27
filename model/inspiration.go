package model

import "gorm.io/gorm"

type Inspiration struct {
	gorm.Model
	BookID              uint   `json:"bookid" form:"bookid"`
	InspirationContexts string `json:"inspirationcontexts" form:"inspirationcontexts" gorm:"type:text"`
}
