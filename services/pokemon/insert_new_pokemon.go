package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

type InsertPokemonService struct {
	PokemonService
	newPokemon Pokemon.Type
}

func NewInsertPokemonService(pokemonService PokemonService, newPokemon Pokemon.Type) IBasePokemonService {
	return &InsertPokemonService{
		PokemonService: pokemonService,
		newPokemon: newPokemon,
	}
}

func getNewPokemonId(pokemons []Pokemon.Type) int {
	lastPokemon := pokemons[len(pokemons) - 1]
	return lastPokemon.ID + 1
}

func (service *InsertPokemonService) Run() (int, Helpers.Response) {
	service.newPokemon.ID = getNewPokemonId(*service.Pokemons)
	*service.Pokemons = append(*service.Pokemons, service.newPokemon)

	return Helpers.SuccessResponse("INSERT_SUCCESSFUL", nil)
}
