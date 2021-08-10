package services

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func TestGetAllPokemonService(t *testing.T) {

	Convey("Get All Pokemon Service", t, func() {

		Convey("Preparation Get All Pokemon Data", func() {
			service := NewGetAllPokemonService(PokemonService{
				Pokemons: Pokemon.GetMockData(),
			})
			expectedHttpCode, expectedResponseData :=
				Helpers.SuccessResponse("DATA_FOUND", *Pokemon.GetMockData())

			Convey("Get All Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})

		Convey("Preparation Get Empty Pokemon Data", func() {
			service := NewGetAllPokemonService(PokemonService{
				Pokemons: &[]Pokemon.Type{},
			})

			expectedHttpCode, expectedResponseData :=
				Helpers.DataNotFoundResponse()

			Convey("Get Empty Pokemon Data", func() {
				httpCode, responseData := service.Run()

				So(httpCode, ShouldEqual, expectedHttpCode)
				So(responseData, ShouldResemble, expectedResponseData)
			})
		})
	})
}
