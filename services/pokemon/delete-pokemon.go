package services

import (
	"net/http"
	"strconv"

	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func DeletePokemon(req *http.Request) (int, Helpers.Response) {
	pokemons := Pokemon.GetData()
	pokemonId, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err == nil {
		for index, value := range *pokemons {
			if value.ID == pokemonId {
				copy((*pokemons)[index:], (*pokemons)[index+1:])
				(*pokemons)[len(*pokemons)-1] = Pokemon.Type{}
				*pokemons = (*pokemons)[:len(*pokemons)-1]

				return Helpers.SuccessResponse("DELETE_SUCCESSFUL", nil)
			}
		}

		return Helpers.DataNotFoundResponse()
	}

	return Helpers.InternalServerErrorResponse()
}