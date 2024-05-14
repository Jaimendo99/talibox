package models

import "gorm.io/gorm"

type GenreCatalog struct {
	gorm.Model
	Genre string
}
