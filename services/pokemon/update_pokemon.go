package services

import (
	"pokemon-go/helpers"
	"pokemon-go/repository"
)

type UpdatePokemonService struct {
	Repository     repository.IBasePokemonRepository
	updatedPokemon repository.Pokemon
}

func NewUpdatePokemonService(
	repository repository.IBasePokemonRepository,
	updatedPokemon repository.Pokemon,
) IBasePokemonService {
	return &UpdatePokemonService{
		Repository:     repository,
		updatedPokemon: updatedPokemon,
	}
}

func (service *UpdatePokemonService) Run() (int, helpers.Response) {
	err := service.Repository.Update(service.updatedPokemon)

	if err == nil {
		return helpers.SuccessResponse("UPDATE_SUCCESSFUL", nil)
	}

	return helpers.DataNotFoundResponse()
}
