package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string
	RealeseId   string         `gorm:"unique"`
	Genres      []GenreCatalog `gorm:"many2many:genre"`
	Year        int
	ReleaseDate string
}

func (m Movie) GetFields() []string {
	return []string{"ID", "Realese ID", "Title", "Genres", "Release Date"}
}

func (m Movie) GetFieldsValues() []string {
	genres := ""
	for _, genre := range m.Genres {
		genres += genre.Genre + ", "
	}
	return []string{
		strconv.FormatUint(uint64(m.ID), 10),
		m.RealeseId,
		m.Title,
		genres,
		m.ReleaseDate,
	}
}

func (m Movie) GetInstanceName() string {
	return "Movie"
}
