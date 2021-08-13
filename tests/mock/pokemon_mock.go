package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"pokemon-go/repository"
)

type PokemonJsonMockRepo struct {
	Mock mock.Mock
}

func (pokemonJsonRepository *PokemonJsonMockRepo) GetAll() ([]repository.Pokemon, error) {
	arguments := pokemonJsonRepository.Mock.Called()
	return arguments[0].([]repository.Pokemon), nil
}

func (pokemonJsonRepository *PokemonJsonMockRepo) GetById(pokemonId int) (repository.Pokemon, error) {
	arguments := pokemonJsonRepository.Mock.Called(pokemonId)

	if arguments.Get(0) == nil {
		return repository.Pokemon{}, errors.New("DATA_NOT_FOUND")
	}

	pokemon := arguments.Get(0).(repository.Pokemon)
	return pokemon, nil
}

func (pokemonJsonRepository *PokemonJsonMockRepo) Insert(newPokemon repository.Pokemon) error {
	arguments := pokemonJsonRepository.Mock.Called(newPokemon)

	if arguments[0] == nil {
		return nil
	}

	return errors.New("EMPTY_DATA")
}

func (pokemonJsonRepository *PokemonJsonMockRepo) Update(updatedPokemon repository.Pokemon) error {
	arguments := pokemonJsonRepository.Mock.Called(updatedPokemon)

	if arguments.Get(0) == nil {
		return errors.New("DATA_NOT_FOUND")
	}

	return nil
}

func (pokemonJsonRepository *PokemonJsonMockRepo) Delete(pokemonId int) error {
	arguments := pokemonJsonRepository.Mock.Called(pokemonId)

	if arguments.Get(0) == nil {
		return errors.New("DATA_NOT_FOUND")
	}

	return nil
}
