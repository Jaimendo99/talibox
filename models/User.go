package models

import (
	"gorm.io/gorm"
)

type UserLogin struct {
	Username string `json:"username" form:"username" gorm:"unique"`
	Password string `json:"password" form:"password"`
}

type UserSignUp struct {
	FullName  string    `json:"fullname" form:"fullname"`
	UserLogin UserLogin `gorm:"embedded"`
}

type User struct {
	gorm.Model
	UserBasic UserSignUp `gorm:"embedded"`
	Movies    []Movie    `gorm:"many2many:tracking"`
	Admin     bool
}
