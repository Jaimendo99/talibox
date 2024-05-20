package uiModels

import "strconv"

type TopGrosses struct {
	Week   string     `json:"week"`
	Movies []MovieTop `json:"movies"`
}

type MovieTop struct {
	MovieName       string  `json:"movieName"`
	DomesticRelease string  `json:"domesticRealese"`
	WeeklyGross     []int32 `json:"weeklyGross"`
}

type MovieTopParsed struct {
	MovieName       string   `json:"movieName"`
	DomesticRealese string   `json:"domesticRealese"`
	WeeklyGross     []string `json:"weeklyGross"`
}

func (m MovieTop) ParseGross() MovieTopParsed {
	weeklyGrossParsed := make([]string, len(m.WeeklyGross))
	for i, gross := range m.WeeklyGross {
		weeklyGrossParsed[i] = "$" + strconv.Itoa(int(gross))
	}
	return MovieTopParsed{
		MovieName:       m.MovieName,
		DomesticRealese: m.DomesticRelease,
		WeeklyGross:     weeklyGrossParsed,
	}
}
