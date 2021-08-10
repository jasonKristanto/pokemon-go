package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

type IBasePokemonService interface {
	Run() (int, Helpers.Response)
}

type PokemonService struct {
	Pokemons *[]Pokemon.Type
}
