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
	baseUrl := "http://192.168.100.14:8000/"
	url := baseUrl + "weeklytop"

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

func GetWeeklyTopGrossWeek(week string, top int) error {

	baseUrl := "http://192.168.100.1:8000/"

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
