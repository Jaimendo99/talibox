package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"talibox/views/uiModels"
)

func GetWeeklyTopGross() (uiModels.TopGrosses, error) {
	baseUrl := "http://localhost:8000/"
	url := baseUrl + "weeklytop/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return uiModels.TopGrosses{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return uiModels.TopGrosses{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return uiModels.TopGrosses{}, err
	}

	var topGrosses uiModels.TopGrosses
	err = json.Unmarshal(body, &topGrosses)
	if err != nil {
		return uiModels.TopGrosses{}, err
	}

	// Print the parsed data to verify
	fmt.Printf("%+v\n", topGrosses)

	// Example of parsing grosses
	for _, movie := range topGrosses.Movies {
		parsed := movie.ParseGross()
		fmt.Printf("%+v\n", parsed)
	}

	return topGrosses, nil
}

func GetHistoricWeeklyTopGross() ([]uiModels.WeeklyTopGross, error) {
	baseUrl := "http://localhost:8000/"
	url := baseUrl + "historicweekly/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []uiModels.WeeklyTopGross{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []uiModels.WeeklyTopGross{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []uiModels.WeeklyTopGross{}, err
	}

	var weeklyTopGross []uiModels.WeeklyTopGross

	err = json.Unmarshal(body, &weeklyTopGross)
	if err != nil {
		return []uiModels.WeeklyTopGross{}, err
	}

	fmt.Printf("%+v\n", weeklyTopGross)

	return weeklyTopGross, nil
}

func SearchMovieByName(titleId string) ([]uiModels.MovieSearch, error) {
	baseUrl := "http://localhost:8000/"
	url := baseUrl + "search/" + titleId

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []uiModels.MovieSearch{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []uiModels.MovieSearch{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []uiModels.MovieSearch{}, err
	}

	var movie []uiModels.MovieSearch

	err = json.Unmarshal(body, &movie)
	if err != nil {
		return []uiModels.MovieSearch{}, err
	}

	fmt.Printf("%+v\n", movie)

	return movie, nil
}

func GetWeeklyByMovie(titleId string) (uiModels.MovieWeeklyGrosses, error) {
	baseUrl := "http://localhost:8000/"
	url := baseUrl + "weeklyhistory/" + titleId

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return uiModels.MovieWeeklyGrosses{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return uiModels.MovieWeeklyGrosses{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return uiModels.MovieWeeklyGrosses{}, err
	}

	var movie uiModels.MovieWeeklyGrosses

	err = json.Unmarshal(body, &movie)
	if err != nil {
		return uiModels.MovieWeeklyGrosses{}, err
	}

	fmt.Printf("%+v\n", movie)

	return movie, nil
}

func GetWeeklyTopGrossWeek(week string, top int) error {

	baseUrl := "http://192.168.100.14:8000/"

	url := baseUrl + "/weeklytop?week=" + week + "&top=" + strconv.FormatInt(int64(top), 10)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
