package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func TestUpdatePokemonService(t *testing.T) {

	Convey("Update Pokemon Service", t, func() {

		Convey("Preparation Update Pokemon Data", func() {
			existedPokemon := Pokemon.Type{
				ID:   1,
				Name: "Bulbasaur",
				Types: []string{
					"Grass",
					"Poison",
				},
				Weaknesses: []string{
					"Fire",
					"Ice",
					"Flying",
					"Psychic",
				},
			}

			service := NewUpdatePokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, existedPokemon)

			expectedHttpCode, expectedResponseData :=
				Helpers.SuccessResponse("UPDATE_SUCCESSFUL", nil)

			Convey("Update Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Not Found Pokemon Data", func() {
			existedPokemon := Pokemon.Type{
				ID:   10,
				Name: "Bulbasaur",
				Types: []string{
					"Grass",
					"Poison",
				},
				Weaknesses: []string{
					"Fire",
					"Ice",
					"Flying",
					"Psychic",
				},
			}

			service := NewUpdatePokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, existedPokemon)

			expectedHttpCode, expectedResponseData :=
				Helpers.DataNotFoundResponse()

			Convey("Not Found Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
