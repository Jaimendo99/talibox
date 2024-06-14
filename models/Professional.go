package models

import "gorm.io/gorm"

type Professional struct {
	gorm.Model
	ArtisticName string `gorm:"unique"`
}
