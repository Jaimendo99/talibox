package models

import (
	"strconv"

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

func (u User) GetFields() []string {
	return []string{"ID", "Full Name", "Username", "Admin"}
}

func (u User) GetFieldsValues() []string {
	return []string{
		strconv.FormatInt(int64(u.ID), 10),
		u.UserBasic.FullName,
		u.UserBasic.UserLogin.Username,
		strconv.FormatBool(u.Admin),
	}
}

func (u User) GetInstanceName() string {
	return "User"
}
