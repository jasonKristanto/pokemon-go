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
	pokemonIndex, err := Pokemon.SearchData(*service.Pokemons, service.updatedPokemon.ID)

	if err == nil {
		(*service.Pokemons)[pokemonIndex] = Pokemon.Type {
			ID: (*service.Pokemons)[pokemonIndex].ID,
			Name: service.updatedPokemon.Name,
			Types: service.updatedPokemon.Types,
			Weaknesses: service.updatedPokemon.Weaknesses,
		}
		return Helpers.SuccessResponse("UPDATE_SUCCESSFUL", nil)
	}

	return Helpers.DataNotFoundResponse()
}
