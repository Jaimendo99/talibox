package uiModels

type MovieSearch struct {
	Title   string   `json:"title"`
	TitleId string   `json:"titleId"`
	Year    string   `json:"year"`
	Actors  []string `json:"actors"`
	ImgUrl  string   `json:"img"`
	Genres  []string `json:"genres"`
}

type MovieWeeklyGrosses struct {
	DomesticRealese string `json:"domesticRealese"`
	SummGross       struct {
		Domestic      int64 `json:"domestic"`
		International int64 `json:"international"`
		Worldwide     int64 `json:"worldwide"`
	} `json:"sumGross"`
	WeeklyGross []struct {
		Rank      int    `json:"rank"`
		Week      int    `json:"week"`
		Gross     int64  `json:"Weekly"`
		WeekStart string `json:"weekStart"`
	} `json:"weekly"`
}
