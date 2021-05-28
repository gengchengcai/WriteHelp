package dao

import (
	"gorm.io/gorm"
	"writer/model"
)

func UserLoginDao(db *gorm.DB, user model.Author) bool {
	newUser := &model.Author{}
	db.Where("username = ?", user.UserName).First(newUser)
	if newUser.ID != 0 {
		return user.PassWord == newUser.PassWord
	}
	return false
}

func UserRegisterDao(db *gorm.DB, user model.Author) bool {
	res := db.Create(&user).RowsAffected
	return res == 1
}
