package services

import (
	"pokemon-go/helpers"
	"pokemon-go/repository"
)

type InsertPokemonService struct {
	Repository repository.IBasePokemonRepository
	newPokemon repository.Pokemon
}

func NewInsertPokemonService(
	repository repository.IBasePokemonRepository, newPokemon repository.Pokemon,
) IBasePokemonService {
	return &InsertPokemonService{
		Repository: repository,
		newPokemon: newPokemon,
	}
}

func (service *InsertPokemonService) Run() (int, helpers.Response) {
	_ = service.Repository.Insert(service.newPokemon)
	return helpers.SuccessResponse("INSERT_SUCCESSFUL", nil)
}
