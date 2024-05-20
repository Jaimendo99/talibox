package models

import (
	"strconv"
	"time"

	"gorm.io/gorm"
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

func (m Movie) GetFields() []string {
	return []string{"ID", "Title", "Genres", "Release Date", "Director"}
}

func (m Movie) GetFieldsValues() []string {
	genres := ""
	for _, genre := range m.Genres {
		genres += genre.Genre + ", "
	}
	return []string{
		strconv.FormatUint(uint64(m.ID), 10),
		m.Title,
		genres,
		m.ReleaseDate.String(),
		m.Director.ArtisticName,
	}
}

func (m Movie) GetInstanceName() string {
	return "Movie"
}
