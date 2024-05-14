package models

import (
	"gorm.io/gorm"
	"time"
)

type WeeklyGrosses struct {
	gorm.Model
	Movie     Movie `gorm:"foreignKey:MovieId"`
	MovieId   int
	weekStart time.Time
	weekEnd   time.Time
	Gross     float64
}
