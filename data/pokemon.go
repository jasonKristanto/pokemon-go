package pokemon

type Type struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Types      []string `json:"types"`
	Weaknesses []string `json:"weaknesses"`
}

var Data = []Type{
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