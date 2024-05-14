package models

import "gorm.io/gorm"

type Professional struct {
	gorm.Model
	ImdbId       string `gorm:"unique"`
	artisticName string
	fullName     string
}
