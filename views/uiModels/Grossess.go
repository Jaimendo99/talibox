package uiModels

import (
	"fmt"
	"strconv"
)

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

type TopGrossesECharts struct {
	Week           string `json:"week"`
	Movies         []string
	GrossFriday    []int32
	GrossSaturday  []int32
	GrossSunday    []int32
	GrossMonday    []int32
	GrossTuesday   []int32
	GrossWednesday []int32
	GrossThursday  []int32
}

func (t TopGrosses) ParseToCharts() TopGrossesECharts {
	movies := make([]string, len(t.Movies))
	grossFriday := make([]int32, len(t.Movies))
	grossSaturday := make([]int32, len(t.Movies))
	grossSunday := make([]int32, len(t.Movies))
	grossMonday := make([]int32, len(t.Movies))
	grossTuesday := make([]int32, len(t.Movies))
	grossWednesday := make([]int32, len(t.Movies))
	grossThursday := make([]int32, len(t.Movies))

	for i, movie := range t.Movies {
		movies[i] = movie.MovieName
		grossFriday[i] = movie.WeeklyGross[0]
		grossSaturday[i] = movie.WeeklyGross[1]
		grossSunday[i] = movie.WeeklyGross[2]
		grossMonday[i] = movie.WeeklyGross[3]
		grossTuesday[i] = movie.WeeklyGross[4]
		grossWednesday[i] = movie.WeeklyGross[5]
		grossThursday[i] = movie.WeeklyGross[6]
	}

	return TopGrossesECharts{
		Week:           t.Week,
		Movies:         movies,
		GrossFriday:    grossFriday,
		GrossSaturday:  grossSaturday,
		GrossSunday:    grossSunday,
		GrossMonday:    grossMonday,
		GrossTuesday:   grossTuesday,
		GrossWednesday: grossWednesday,
		GrossThursday:  grossThursday,
	}
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

type WeeklyTopGross struct {
	Year         string           `json:"year"`
	Week         string           `json:"week"`
	TotalGross   int64            `json:"total_gross"`
	OthersGross  int64            `json:"others_gross"`
	Top5Grossers []MovieGrossWeek `json:"top5_gross"`
}

type MovieGrossWeek struct {
	Rank    int    `json:"Rank"`
	Release string `json:"Release"`
	Gross   int64  `json:"Gross"`
}

type YearWeeklyGross struct {
	Year    string     `json:"year"`
	Grosses []int      `json:"grosses"`
	Movies  [][]string `json:"movies"`
}

func TransformWeeklyToYearlyGross(weeklyData []WeeklyTopGross) []YearWeeklyGross {
	yearGrossMap := make(map[string]YearWeeklyGross)

	// Aggregate the total gross by year
	for _, weeklyRecord := range weeklyData {
		year := weeklyRecord.Year
		totalGross := int(weeklyRecord.TotalGross)

		// Convert top 5 movies to a list of strings
		var topMovies []string
		for _, movie := range weeklyRecord.Top5Grossers {
			topMovies = append(topMovies, fmt.Sprintf("%d. %s: $%d", movie.Rank, movie.Release, movie.Gross))
		}

		if yg, exists := yearGrossMap[year]; exists {
			yg.Grosses = append(yg.Grosses, totalGross)
			yg.Movies = append(yg.Movies, topMovies)
			yearGrossMap[year] = yg
		} else {
			yearGrossMap[year] = YearWeeklyGross{
				Year:    year,
				Grosses: []int{totalGross},
				Movies:  [][]string{topMovies},
			}
		}
	}

	// Convert the map to a slice of YearWeeklyGross
	var yearlyGrossList []YearWeeklyGross
	for _, yg := range yearGrossMap {
		yearlyGrossList = append(yearlyGrossList, yg)
	}

	return yearlyGrossList
}
