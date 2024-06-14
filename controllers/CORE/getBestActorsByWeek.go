package core

import (
	models1 "talibox/controllers/Models"
	"talibox/inits"
	"talibox/models"
)

func GetBestActorsByWeek(weeks []int) []models1.ActorGross {
	actorsGross := []models1.ActorGross{}

	for _, week := range weeks {
		actorsGross = append(actorsGross, GetActorsAndGrossByAWeek(week)...)
	}

	return actorsGross
}

func GetActorsAndGrossByAWeek(week int) []models1.ActorGross {
	db := inits.DB

	//gross by week
	grosses := []models.WeeklyGrosses{}
	db.Find(&grosses)

	grossesByWeek := []models.WeeklyGrosses{}
	for _, gross := range grosses {
		if gross.Week == week {
			grossesByWeek = append(grossesByWeek, gross)
		}
	}
	//get movies from those grosses
	movies := []models.Movie{}
	db.Find(&movies)

	moviesByWeek := []models.Movie{}
	for _, movie := range movies {
		for _, gross := range grossesByWeek {
			if movie.ID == uint(gross.MovieId) {
				moviesByWeek = append(moviesByWeek, movie)
			}
		}
	}

	//get actors from those movies cast connects movie to professional
	actors := []models.Professional{}
	db.Find(&actors)

	casts := []models.Cast{}
	db.Find(&casts)

	actorsByWeek := []models.Professional{}
	for _, actor := range actors {
		for _, cast := range casts {
			for _, movie := range moviesByWeek {
				if cast.ProfessionalId == actor.ID && cast.MovieId == movie.ID {
					actorsByWeek = append(actorsByWeek, actor)
				}
			}
		}
	}

	//get gross from those actors
	actorsGross := []models1.ActorGross{}
	for _, actor := range actorsByWeek {
		for _, cast := range casts {
			for _, gross := range grossesByWeek {
				if cast.ProfessionalId == actor.ID {
					actorsGross = append(actorsGross, models1.ActorGross{
						Actor: actor.ArtisticName,
						Gross: gross.Gross,
					})
				}
			}
		}
	}

	return actorsGross

}
