package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func TestGetPokemonService(t *testing.T) {

	Convey("Get Pokemon Service", t, func() {

		Convey("Preparation Get Pokemon Data", func() {
			searchedPokemonId := 2
			searchedPokemonIndex, _ := Pokemon.SearchData(*Pokemon.GetMockData(), searchedPokemonId)

			service := NewGetPokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, searchedPokemonId)

			expectedHttpCode, expectedResponseData :=
				Helpers.SuccessResponse("DATA_FOUND", (*Pokemon.GetMockData())[searchedPokemonIndex])

			Convey("Get Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Not Found Pokemon Data", func() {
			searchedPokemonId := 10

			service := NewGetPokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			}, searchedPokemonId)

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