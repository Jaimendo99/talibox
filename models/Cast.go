package models

type Cast struct {
	MovieId        uint `gorm:"primaryKey"`
	ProfessionalId uint `gorm:"primaryKey"`

	Movie        Movie        `gorm:"foreignKey:MovieId"`
	Professional Professional `gorm:"primaryKey"`
	Role         string
}
