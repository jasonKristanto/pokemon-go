package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func TestInsertPokemonService(t *testing.T) {

	Convey("Insert Pokemon Service", t, func() {

		Convey("Preparation Insert Pokemon Data", func() {
			newPokemon := Pokemon.Type{
				ID:   0,
				Name: "Chameleon",
				Types: []string{
					"Fire",
					"Attack",
					"Flying",
				},
				Weaknesses: []string{
					"Water",
					"Grass",
				},
			}

			service := NewInsertPokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, newPokemon)

			expectedHttpCode, expectedResponseData :=
				Helpers.SuccessResponse("INSERT_SUCCESSFUL", nil)

			Convey("Insert Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
