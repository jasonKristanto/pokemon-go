package services

import (
	"pokemon-go/helpers"
	"pokemon-go/repository"
)

type GetPokemonService struct {
	Repository repository.IBasePokemonRepository
	pokemonId int
}

func NewGetPokemonService(repository repository.IBasePokemonRepository, pokemonId int) IBasePokemonService {
	return &GetPokemonService{
		Repository: repository,
		pokemonId: pokemonId,
	}
}

func (service *GetPokemonService) Run() (int, helpers.Response) {
	pokemon, err := service.Repository.GetById(service.pokemonId)

	if err == nil {
		return helpers.SuccessResponse("DATA_FOUND", pokemon)
	}

	return helpers.DataNotFoundResponse()
}
