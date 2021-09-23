package services

import (
	"pokemon-go/helpers"
	"pokemon-go/repository"
)

type DeletePokemonService struct {
	Repository repository.IBasePokemonRepository
	pokemonId  int
}

func NewDeletePokemonService(repository repository.IBasePokemonRepository, pokemonId int) IBasePokemonService {
	return &DeletePokemonService{
		Repository: repository,
		pokemonId:  pokemonId,
	}
}

func (service *DeletePokemonService) Run() (int, helpers.Response) {
	err := service.Repository.Delete(service.pokemonId)

	if err == nil {
		return helpers.SuccessResponse("DELETE_SUCCESSFUL", nil)
	}

	return helpers.DataNotFoundResponse()
}
