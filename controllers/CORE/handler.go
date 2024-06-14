package core

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"talibox/controllers/Models"
	"talibox/views/actors"
)

func GetActorsAndGrossByAWeekNet() echo.HandlerFunc {
	return func(c echo.Context) error {
		weekStart := c.FormValue("weekStart")
		weekEnd := c.FormValue("weekEnd")

		weeksint, err := strconv.ParseInt(weekStart, 10, 64)
		if err != nil {
			return c.JSON(400, "weekStart must be an integer")
		}
		weekEndint, err := strconv.ParseInt(weekEnd, 10, 64)
		if err != nil {
			return c.JSON(400, "weekEnd must be an integer")
		}

		weeks := []int{}

		for i := weeksint; i <= weekEndint; i++ {
			weeks = append(weeks, int(i))
		}

		actorsGross := GetBestActorsByWeek(weeks)
		aggregatedGross := AggregateActorGross(actorsGross)

		return actors.ActorGrossView(aggregatedGross).Render(c.Request().Context(), c.Response().Writer)

	}
}

func AggregateActorGross(actorsGross []Models.ActorGross) []Models.ActorGross {
	aggregatedMap := make(map[string]int64)
	for _, ag := range actorsGross {
		aggregatedMap[ag.Actor] += ag.Gross
	}

	aggregated := []Models.ActorGross{}
	for actor, gross := range aggregatedMap {
		aggregated = append(aggregated, Models.ActorGross{Actor: actor, Gross: gross})
	}
	return aggregated
}

func GetActorForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		return actors.ActorInputs().Render(c.Request().Context(), c.Response().Writer)
	}
}
