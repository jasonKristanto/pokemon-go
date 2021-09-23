package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"pokemon-go/helpers"
	"pokemon-go/repository"
)

func TestDeletePokemonService(t *testing.T) {

	Convey("Delete Pokemon Service", t, func() {

		Convey("Preparation Delete Pokemon Data", func() {
			deletedPokemon := repository.Pokemon{
				ID:         5,
				Name:       "Kakuna",
				Types:      []string{"Bug", "Poison"},
				Weaknesses: []string{"Fire", "Flying", "Psychic", "Rock"},
			}
			pokemonRepositoryMock.Mock.On("Delete", deletedPokemon.ID).Return(deletedPokemon)

			service := NewDeletePokemonService(pokemonRepositoryMock, deletedPokemon.ID)

			expectedHttpCode, expectedResponseData :=
				helpers.SuccessResponse("DELETE_SUCCESSFUL", nil)

			Convey("Delete Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Not Found Pokemon Data", func() {
			notFoundPokemonId := 10
			pokemonRepositoryMock.Mock.On("Delete", notFoundPokemonId).Return(nil)

			service := NewDeletePokemonService(pokemonRepositoryMock, notFoundPokemonId)

			expectedHttpCode, expectedResponseData :=
				helpers.DataNotFoundResponse()

			Convey("Found Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
