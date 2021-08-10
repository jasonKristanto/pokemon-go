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
	for index, value := range *service.Pokemons {
		if value.ID == service.pokemonId {
			copy((*service.Pokemons)[index:], (*service.Pokemons)[index+1:])
			(*service.Pokemons)[len(*service.Pokemons)-1] = Pokemon.Type{}
			*service.Pokemons = (*service.Pokemons)[:len(*service.Pokemons)-1]

			return Helpers.SuccessResponse("DELETE_SUCCESSFUL", nil)
		}
	}

	return Helpers.DataNotFoundResponse()
}