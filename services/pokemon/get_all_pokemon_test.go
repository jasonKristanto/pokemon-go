package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"pokemon-go/helpers"
	"pokemon-go/repository"
)

func TestGetAllPokemonService(t *testing.T) {

	Convey("Get All Pokemon Service", t, func() {

		Convey("Preparation Get All Pokemon Data", func() {
			allPokemons := []repository.Pokemon{
				{
					1,
					"Ivysaur",
					[]string{"Grass", "Poison"},
					[]string{"Fire", "Ice", "Flying", "Psychic"},
				},
				{
					2,
					"Charmeleon",
					[]string{"Fire"},
					[]string{"Water", "Ground", "Rock"},
				},
				{
					3,
					"Wartortle",
					[]string{"Water"},
					[]string{"Electric", "Grass"},
				},
				{
					4,
					"Metapod",
					[]string{"Bug"},
					[]string{"Fire", "Flying", "Rock"},
				},
				{
					5,
					"Kakuna",
					[]string{"Bug", "Poison"},
					[]string{"Fire", "Flying", "Psychic", "Rock"},
				},
			}
			pokemonRepositoryMock.Mock.On("GetAll").Return(allPokemons)

			service := NewGetAllPokemonService(pokemonRepositoryMock)

			expectedHttpCode, expectedResponseData :=
				helpers.SuccessResponse("DATA_FOUND", allPokemons)

			Convey("Get All Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
