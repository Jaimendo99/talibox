package models

import "gorm.io/gorm"

type Professional struct {
	gorm.Model
	ImdbId       string `gorm:"unique"`
	ArtisticName string
	FullName     string
}
