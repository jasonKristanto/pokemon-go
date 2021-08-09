package services

import (
	"net/http"
	"strconv"

	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func GetPokemon(req *http.Request) (int, Helpers.Response) {
	pokemons := Pokemon.GetData()
	pokemonId, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err == nil {
		for _, value := range *pokemons {
			if value.ID == pokemonId {
				return Helpers.SuccessResponse("DATA_FOUND", value)
			}
		}

		return Helpers.DataNotFoundResponse()
	}

	return Helpers.InternalServerErrorResponse()
}
