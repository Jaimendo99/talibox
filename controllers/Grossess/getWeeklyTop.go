package grossess

import (
	external "talibox/controllers/External"
	homepagecontroller "talibox/controllers/HomePageController"
	groses "talibox/views/Grosess"
	"talibox/views/uiModels"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/labstack/echo/v4"
)

func GetWeeklyTop() echo.HandlerFunc {
	return func(c echo.Context) error {
		topGrossess := uiModels.TopGrosses{}

		topGrossess, err := external.GetWeeklyTopGross()

		if err != nil {
			return err
		}

		moviesTop := topGrossess.Movies

		topGrossessParsed := []uiModels.MovieTopParsed{}

		for _, movie := range moviesTop {
			topGrossessParsed = append(topGrossessParsed, movie.ParseGross())
		}

		component := groses.TopGrossList(topGrossessParsed)
		header := c.Request().Header.Get("Hx-Request")

		if header == "" {
			component = homepagecontroller.PageDefaultContainer(false, component, "Weekly Top Gross", "Guest")
		}

		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}

func genBarItems(groses []int32) []opts.BarData {
	var items []opts.BarData
	for _, gross := range groses {
		items = append(items, opts.BarData{Value: gross})
	}
	return items

}

func GrosessChart(groses uiModels.TopGrossesECharts) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Top Grosses week: " + groses.Week,
			Top:   "-4",
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
			Top:  "30",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
			Name:  "Total Gross",
			AxisLabel: &opts.AxisLabel{
				Show:      true,
				Formatter: "${value}",
			},
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Formatter: "{a} <br/>{b} : ${c}",
		}),
		charts.WithColorsOpts(opts.Colors{
			// "#577590", "#43aa8b", "#90be6d", "#f9c74f", "#f8961e", "#f3722c", "#f94144",
			"#264653ff", "#287271ff", "#2a9d8fff", "#8ab17dff", "#e9c46aff", "#f4a261ff", "#e76f51ff",
		}),
	)
	bar.SetXAxis(groses.Movies).
		AddSeries("Friday", genBarItems(groses.GrossFriday)).
		AddSeries("Saturday", genBarItems(groses.GrossSaturday)).
		AddSeries("Sunday", genBarItems(groses.GrossSunday)).
		AddSeries("Monday", genBarItems(groses.GrossMonday)).
		AddSeries("Tuesday", genBarItems(groses.GrossTuesday)).
		AddSeries("Wednesday", genBarItems(groses.GrossWednesday)).
		AddSeries("Thursday", genBarItems(groses.GrossThursday)).
		SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
			Stack: "gross",
		}))

	return bar
}

func GetWeeklyTopChart() echo.HandlerFunc {
	return func(c echo.Context) error {

		htmxHd := c.Request().Header.Get("Hx-Request")

		topGrossess, err := external.GetWeeklyTopGross()
		if err != nil {
			return err
		}

		topGrossessECharts := topGrossess.ParseToCharts()

		bar := GrosessChart(topGrossessECharts)

		if htmxHd == "" {
			return homepagecontroller.PageDefaultContainer(
				false,
				groses.GetChart(),
				"Weekly Top Gross",
				"Guest",
			).Render(c.Request().Context(), c.Response().Writer)
		}

		return bar.Render(c.Response().Writer)
	}
}
