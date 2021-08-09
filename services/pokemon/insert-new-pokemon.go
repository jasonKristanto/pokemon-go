package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func getNewPokemonId(pokemons []Pokemon.Type) int {
	lastPokemon := pokemons[len(pokemons) - 1]
	return lastPokemon.ID + 1
}

func InsertNewPokemon(newPokemon Pokemon.Type) (int, Helpers.Response) {
	pokemons := Pokemon.GetData()
	newPokemon.ID = getNewPokemonId(*pokemons)
	*pokemons = append(*pokemons, newPokemon)

	return Helpers.SuccessResponse("INSERT_SUCCESSFUL", nil)
}
