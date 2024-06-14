package grossess

import (
	"math"
	external "talibox/controllers/External"
	"talibox/controllers/Models"
	"talibox/views/Insights"
	"talibox/views/uiModels"

	"github.com/labstack/echo/v4"
)

func GetYearToYearComparason() echo.HandlerFunc {
	return func(c echo.Context) error {
		variances := GetVariance()
		standardDeviations := GetStandardDeviations()

		return Insights.ShowInsights(variances, standardDeviations).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetVariance() []Models.Variances {

	weeklyYear, err := external.GetHistoricWeeklyTopGross()

	if err != nil {
		return []Models.Variances{}
	}

	yearWeeklyGross := uiModels.TransformWeeklyToYearlyGross(weeklyYear)

	var variances []Models.Variances

	for _, yearGross := range yearWeeklyGross {
		// get the variance
		year := yearGross.Year
		var weekSum int
		for _, week := range yearGross.Grosses {
			weekSum += week
		}

		avg := float32(weekSum) / float32(len(yearGross.Grosses))

		var iMenusAvg float32

		for _, week := range yearGross.Grosses {
			iMenusAvg += (float32(week) - avg) * (float32(week) - avg)
		}
		variance := iMenusAvg / float32((len(yearGross.Grosses) - 1))

		variances = append(variances, Models.Variances{
			Year:     year,
			Variance: variance,
		})

	}

	return variances
}

func GetStandardDeviations() []Models.StandardDeviations {

	vaiances := GetVariance()

	var standardDeviations []Models.StandardDeviations

	for _, variance := range vaiances {
		standardDeviations = append(standardDeviations, Models.StandardDeviations{
			Year:              variance.Year,
			StandardDeviation: float32(math.Sqrt(float64(variance.Variance))),
		})
	}

	return standardDeviations
}
