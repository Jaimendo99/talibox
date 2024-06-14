package models

import (
	"gorm.io/gorm"
)

type WeeklyGrosses struct {
	gorm.Model
	Movie     Movie `gorm:"foreignKey:MovieId"`
	MovieId   int
	Rank      int
	WeekStart string
	Week      int
	Gross     int64
}
