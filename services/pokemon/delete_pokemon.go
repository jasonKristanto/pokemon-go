package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

type DeletePokemonService struct {
	PokemonService
	pokemonId int
}

func NewDeletePokemonService(pokemonService PokemonService, pokemonId int) IBasePokemonService {
	return &DeletePokemonService{
		PokemonService: pokemonService,
		pokemonId: pokemonId,
	}
}

func (service *DeletePokemonService) Run() (int, Helpers.Response) {
	pokemonIndex, err := Pokemon.SearchData(*service.Pokemons, service.pokemonId)

	if err == nil {
		copy((*service.Pokemons)[pokemonIndex:], (*service.Pokemons)[pokemonIndex+1:])
		(*service.Pokemons)[len(*service.Pokemons)-1] = Pokemon.Type{}
		*service.Pokemons = (*service.Pokemons)[:len(*service.Pokemons)-1]

		return Helpers.SuccessResponse("DELETE_SUCCESSFUL", nil)
	}

	return Helpers.DataNotFoundResponse()
}