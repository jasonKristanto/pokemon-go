package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func UpdatePokemon(updatedPokemon Pokemon.Type) (int, Helpers.Response) {
	pokemons := Pokemon.GetData()

	for index, value := range *pokemons {
		if value.ID == updatedPokemon.ID {
			(*pokemons)[index] = Pokemon.Type {
				ID: value.ID,
				Name: updatedPokemon.Name,
				Types: updatedPokemon.Types,
				Weaknesses: updatedPokemon.Weaknesses,
			}
			return Helpers.SuccessResponse("UPDATE_SUCCESSFUL", nil)
		}
	}

	return Helpers.DataNotFoundResponse()
}
