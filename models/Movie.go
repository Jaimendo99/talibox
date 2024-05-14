package models

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	ImdbId      string `gorm:"unique"`
	Title       string
	Genres      []GenreCatalog `gorm:"many2many:genre"`
	ReleaseDate time.Time
	DirectorId  uint
	Director    Professional `gorm:"foreignKey:DirectorId"`
}
