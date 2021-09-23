package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"pokemon-go/helpers"
	"pokemon-go/repository"
)

func TestInsertPokemonService(t *testing.T) {
	Convey("Insert Pokemon Service", t, func() {

		Convey("Preparation Insert Pokemon Data", func() {
			newPokemon := repository.Pokemon{
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
			pokemonRepositoryMock.Mock.On("Insert", newPokemon).Return(nil)

			service := NewInsertPokemonService(pokemonRepositoryMock, newPokemon)

			expectedHttpCode, expectedResponseData :=
				helpers.SuccessResponse("INSERT_SUCCESSFUL", nil)

			Convey("Insert Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
