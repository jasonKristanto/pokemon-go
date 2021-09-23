package repository

import (
	"errors"
	"sync"
)

var pokemonData *[]Pokemon

type PokemonJsonRepo struct {
}

func init() {
	if pokemonData == nil {
		var syncOnce sync.Once
		syncOnce.Do(func() {
			pokemonData = &[]Pokemon{
				{
					1,
					"Bulbasaur",
					[]string{"Grass", "Poison"},
					[]string{"Fire", "Ice", "Flying", "Psychic"},
				},
				{
					2,
					"Charmander",
					[]string{"Fire"},
					[]string{"Water", "Ground", "Rock"},
				},
				{
					3,
					"Squirtle",
					[]string{"Water"},
					[]string{"Electric", "Grass"},
				},
				{
					4,
					"Caterpie",
					[]string{"Bug"},
					[]string{"Fire", "Flying", "Rock"},
				},
				{
					5,
					"Weedle",
					[]string{"Bug", "Poison"},
					[]string{"Fire", "Flying", "Psychic", "Rock"},
				},
			}
		})
	}
}

func (pokemonJsonRepository *PokemonJsonRepo) GetAll() ([]Pokemon, error) {
	if pokemonData == nil {
		return nil, errors.New("DATA_NOT_FOUND")
	}

	return *pokemonData, nil
}

func (pokemonJsonRepository *PokemonJsonRepo) GetById(pokemonId int) (Pokemon, error) {
	for _, value := range *pokemonData {
		if value.ID == pokemonId {
			return value, nil
		}
	}

	return Pokemon{}, errors.New("DATA_NOT_FOUND")
}

func getNewPokemonId() int {
	lastPokemon := (*pokemonData)[len(*pokemonData)-1]
	return lastPokemon.ID + 1
}

func (pokemonJsonRepository *PokemonJsonRepo) Insert(newPokemon Pokemon) error {
	newPokemon.ID = getNewPokemonId()
	*pokemonData = append(*pokemonData, newPokemon)

	return nil
}

func (pokemonJsonRepository *PokemonJsonRepo) Update(updatedPokemon Pokemon) error {
	for index, value := range *pokemonData {
		if value.ID == updatedPokemon.ID {
			(*pokemonData)[index] = Pokemon{
				ID:         (*pokemonData)[index].ID,
				Name:       updatedPokemon.Name,
				Types:      updatedPokemon.Types,
				Weaknesses: updatedPokemon.Weaknesses,
			}

			return nil
		}
	}

	return errors.New("DATA_NOT_FOUND")
}

func (pokemonJsonRepository *PokemonJsonRepo) Delete(pokemonId int) error {
	for index, value := range *pokemonData {
		if value.ID == pokemonId {
			copy((*pokemonData)[index:], (*pokemonData)[index+1:])
			(*pokemonData)[len(*pokemonData)-1] = Pokemon{}
			*pokemonData = (*pokemonData)[:len(*pokemonData)-1]

			return nil
		}
	}

	return errors.New("DATA_NOT_FOUND")
}
