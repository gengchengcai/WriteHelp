package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	CreditCards []CreditCard
	Tests       []Test
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type Test struct {
	UserID uint
}
