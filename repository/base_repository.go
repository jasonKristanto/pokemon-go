package repository

type Pokemon struct {
	ID         int      `json:"id" pg:",pk"`
	Name       string   `json:"name"`
	Types      []string `json:"types" pg:",array"`
	Weaknesses []string `json:"weaknesses" pg:",array"`
}

type IBasePokemonRepository interface {
	GetAll() ([]Pokemon, error)
	GetById(pokemonId int) (Pokemon, error)
	Insert(newPokemon Pokemon) error
	Update(updatedPokemon Pokemon) error
	Delete(pokemonId int) error
}
