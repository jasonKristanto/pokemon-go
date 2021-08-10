package data

import (
	"sync"
)

var PokemonMockData *[]Type

func GetMockData() *[]Type {
	if PokemonMockData == nil {
		var syncOnce sync.Once
		syncOnce.Do(func() {
			PokemonMockData = &[]Type{
				{
					1,
					"Ivysaur",
					[]string{"Grass", "Poison"},
					[]string{"Fire", "Ice", "Flying", "Psychic"},
				},
				{
					2,
					"Charmeleon",
					[]string{"Fire"},
					[]string{"Water", "Ground", "Rock"},
				},
				{
					3,
					"Wartortle",
					[]string{"Water"},
					[]string{"Electric", "Grass"},
				},
				{
					4,
					"Metapod",
					[]string{"Bug"},
					[]string{"Fire", "Flying", "Rock"},
				},
				{
					5,
					"Kakuna",
					[]string{"Bug", "Poison"},
					[]string{"Fire", "Flying", "Psychic", "Rock"},
				},
			}
		})
	}

	return PokemonMockData
}
