package model

import "gorm.io/gorm"

type Rolesetting struct {
	gorm.Model
	RolesettingContexts string `json:"rolesettingcontexts" form:"rolesettingcontexts" gorm:"type:text"`
	BookID              uint   `json:"bookid" form:"bookid"`
}
