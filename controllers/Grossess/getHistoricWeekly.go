package grossess

import (
	"strings"
	external "talibox/controllers/External"
	homepagecontroller "talibox/controllers/HomePageController"
	"talibox/views/Grosess"
	"talibox/views/uiModels"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/labstack/echo/v4"
)

func GetHistoricWeekly() echo.HandlerFunc {
	return func(c echo.Context) error {

		weeklyTopGross, err := external.GetHistoricWeeklyTopGross()
		if err != nil {
			return err
		}

		return Grosess.HistoricGrossess(weeklyTopGross).Render(c.Request().Context(), c.Response().Writer)

	}
}

func genLineItems(yearGross uiModels.YearWeeklyGross) []opts.LineData {
	var items []opts.LineData
	for i, value := range yearGross.Grosses {
		// Create the movie list as a single string
		movieList := "No data"
		if i < len(yearGross.Movies) {
			movieList = joinMovies(yearGross.Movies[i])
		}
		items = append(items, opts.LineData{
			Name:  movieList,
			Value: value,
		})
	}
	return items
}

func joinMovies(movies []string) string {
	return strings.Join(movies, "\n")
}

func HistoricGrossessChart(grosses []uiModels.YearWeeklyGross) *charts.Line {
	weeks := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52",
	}

	c := charts.NewLine()

	c.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title: "Historic Yearly Gross per week",
				Top:   "5%",
			},
		),
		charts.WithDataZoomOpts(
			opts.DataZoom{
				Type:       "slider",
				Start:      0,
				End:        100,
				XAxisIndex: []int{0},
			},
		),
		charts.WithToolboxOpts(
			opts.Toolbox{
				Show:  true,
				Right: "20%",
				Feature: &opts.ToolBoxFeature{
					SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
						Show:  true,
						Title: "Save as image",
					},
				},
			},
		),
		charts.WithTooltipOpts(
			opts.Tooltip{
				Show:    true,
				Trigger: "item",
				AxisPointer: &opts.AxisPointer{
					Type: "cross",
				},
				Formatter: "<div>{a}</div><div  style=' white-space: pre; text-overflow: ellipsis;'>{b}</div><br><div>Total Gross: ${c}</div>",
			},
		),
	)

	c.SetXAxis(weeks)

	for _, yearGross := range grosses {
		c.AddSeries(yearGross.Year, genLineItems(yearGross),
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: false,
			}),
		)
	}

	return c
}

func GetHistoricYearWeekly() echo.HandlerFunc {
	return func(c echo.Context) error {
		htmxHd := c.Request().Header.Get("Hx-Request")

		yearWeeklyGroses, err := external.GetHistoricWeeklyTopGross()

		if err != nil {
			return err
		}

		data := uiModels.TransformWeeklyToYearlyGross(yearWeeklyGroses)

		line := HistoricGrossessChart(data)

		if htmxHd == "" {
			return homepagecontroller.PageDefaultContainer(
				false,
				Grosess.GetCoreData(),
				"Historic Yearly Gross per week",
				"Guest",
			).Render(c.Request().Context(), c.Response().Writer)
		}

		return line.Render(c.Response().Writer)
	}
}
