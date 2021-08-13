package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"pokemon-go/helpers"
	"pokemon-go/repository"
)

func TestGetPokemonService(t *testing.T) {
	Convey("Get Pokemon Service", t, func() {

		Convey("Preparation Get Pokemon Data", func() {
			searchedPokemon := repository.Pokemon{
				ID:         2,
				Name:       "Charmeleon",
				Types:      []string{"Fire"},
				Weaknesses: []string{"Water", "Ground", "Rock"},
			}
			pokemonRepositoryMock.Mock.On("GetById", 2).Return(searchedPokemon)

			service := NewGetPokemonService(pokemonRepositoryMock, searchedPokemon.ID)

			expectedHttpCode, expectedResponseData :=
				helpers.SuccessResponse("DATA_FOUND", searchedPokemon)

			Convey("Get Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Not Found Pokemon Data", func() {
			notFoundPokemonId := 10
			pokemonRepositoryMock.Mock.On("GetById", notFoundPokemonId).Return(nil)

			service := NewGetPokemonService(pokemonRepositoryMock, notFoundPokemonId)

			expectedHttpCode, expectedResponseData :=
				helpers.DataNotFoundResponse()

			Convey("Get Not Found Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
