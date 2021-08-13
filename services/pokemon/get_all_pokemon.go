package services

import (
	"pokemon-go/helpers"
	"pokemon-go/repository"
)

type GetAllPokemonService struct {
	Repository repository.IBasePokemonRepository
}

func NewGetAllPokemonService(repository repository.IBasePokemonRepository) IBasePokemonService {
	return &GetAllPokemonService{
		Repository: repository,
	}
}

func (service *GetAllPokemonService) Run() (int, helpers.Response) {
	pokemons, err := service.Repository.GetAll()

	if err == nil {
		return helpers.SuccessResponse("DATA_FOUND", pokemons)
	}

	return helpers.DataNotFoundResponse()
}
