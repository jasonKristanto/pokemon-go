package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func TestDeletePokemonService(t *testing.T) {

	Convey("Delete Pokemon Service", t, func() {

		Convey("Preparation Delete Pokemon Data", func() {
			deletedPokemonId := 3

			service := NewDeletePokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, deletedPokemonId)

			expectedHttpCode, expectedResponseData :=
				Helpers.SuccessResponse("DELETE_SUCCESSFUL", nil)

			Convey("Get Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Not Found Pokemon Data", func() {
			deletedPokemonId := 10

			service := NewDeletePokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, deletedPokemonId)

			expectedHttpCode, expectedResponseData :=
				Helpers.DataNotFoundResponse()

			Convey("Get Not Found Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}