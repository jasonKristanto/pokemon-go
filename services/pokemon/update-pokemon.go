package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

type UpdatePokemonService struct {
	PokemonService
	updatedPokemon Pokemon.Type
}

func NewUpdatePokemonService(pokemonService PokemonService, updatedPokemon Pokemon.Type) IBasePokemonService {
	return &UpdatePokemonService{
		PokemonService: pokemonService,
		updatedPokemon: updatedPokemon,
	}
}

func (service *UpdatePokemonService) Run() (int, Helpers.Response) {
	for index, value := range *service.Pokemons {
		if value.ID == service.updatedPokemon.ID {
			(*service.Pokemons)[index] = Pokemon.Type {
				ID: value.ID,
				Name: service.updatedPokemon.Name,
				Types: service.updatedPokemon.Types,
				Weaknesses: service.updatedPokemon.Weaknesses,
			}
			return Helpers.SuccessResponse("UPDATE_SUCCESSFUL", nil)
		}
	}

	return Helpers.DataNotFoundResponse()
}
