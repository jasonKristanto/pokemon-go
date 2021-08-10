package services

import (
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
	for _, value := range *service.Pokemons {
		if value.ID == service.pokemonId {
			return Helpers.SuccessResponse("DATA_FOUND", value)
		}
	}

	return Helpers.DataNotFoundResponse()
}
