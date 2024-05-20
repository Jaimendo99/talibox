package models

import (
	"time"

	"gorm.io/gorm"
)

type WeeklyGrosses struct {
	gorm.Model
	Movie     Movie `gorm:"foreignKey:MovieId"`
	MovieId   int
	WeekStart time.Time
	WeekEnd   time.Time
	Gross     float64
}
