package services

import (
	Helpers "pokemon-go/helpers"
)

type GetAllPokemonService struct {
	PokemonService
}

func NewGetAllPokemonService(pokemonService PokemonService) IBasePokemonService {
	return &GetAllPokemonService{
		PokemonService: pokemonService,
	}
}

func (service *GetAllPokemonService) Run() (int, Helpers.Response) {
	if len(*service.Pokemons) > 0 {
		return Helpers.SuccessResponse("DATA_FOUND", *service.Pokemons)
	}

	return Helpers.DataNotFoundResponse()
}