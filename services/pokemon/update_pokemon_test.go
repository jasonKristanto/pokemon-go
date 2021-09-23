package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"pokemon-go/helpers"
	"pokemon-go/repository"
)

func TestUpdatePokemonService(t *testing.T) {
	Convey("Update Pokemon Service", t, func() {

		Convey("Preparation Update Pokemon Data", func() {
			existedPokemon := repository.Pokemon{
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
			pokemonRepositoryMock.Mock.On("Update", existedPokemon).Return(existedPokemon)

			service := NewUpdatePokemonService(pokemonRepositoryMock, existedPokemon)

			expectedHttpCode, expectedResponseData :=
				helpers.SuccessResponse("UPDATE_SUCCESSFUL", nil)

			Convey("Update Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Not Found Pokemon Data", func() {
			notFoundPokemon := repository.Pokemon{
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
			pokemonRepositoryMock.Mock.On("Update", notFoundPokemon).Return(nil)

			service := NewUpdatePokemonService(pokemonRepositoryMock, notFoundPokemon)

			expectedHttpCode, expectedResponseData :=
				helpers.DataNotFoundResponse()

			Convey("Not Found Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
