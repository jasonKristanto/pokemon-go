package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

type GetPokemonService struct {
	PokemonService
	pokemonId int
}

func NewGetPokemonService(pokemonService PokemonService, pokemonId int) IBasePokemonService {
	return &GetPokemonService{
		PokemonService: pokemonService,
		pokemonId: pokemonId,
	}
}

func (service *GetPokemonService) Run() (int, Helpers.Response) {
	pokemonIndex, err := Pokemon.SearchData(*service.Pokemons, service.pokemonId)

	if err == nil {
		return Helpers.SuccessResponse("DATA_FOUND", (*service.Pokemons)[pokemonIndex])
	}

	return Helpers.DataNotFoundResponse()
}
