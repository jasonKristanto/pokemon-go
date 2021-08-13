package database

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-pg/pg/v10"

	"pokemon-go/config"
	"pokemon-go/repository"
)

type PokemonPgSqlRepo struct {
}

var pgDb *pg.DB

func init() {
	if pgDb == nil {
		var syncOnce sync.Once
		syncOnce.Do(func() {
			pgConfig := config.Configuration.Database

			pgDb = pg.Connect(&pg.Options{
				Addr:     fmt.Sprintf("%s:%d", pgConfig.Host, pgConfig.Port),
				User:     pgConfig.Username,
				Password: pgConfig.Password,
				Database: pgConfig.Database,
			})

			ctx := context.Background()
			if err := pgDb.Ping(ctx); err != nil {
				panic(err)
			}
		})
	}
}

func (pokemonPgSqlRepo *PokemonPgSqlRepo) GetAll() ([]repository.Pokemon, error) {
	var pokemons []repository.Pokemon

	err := pgDb.Model(&pokemons).Order("id").Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("INTERNAL_SERVER_ERROR")
	}

	return pokemons, nil
}

func (pokemonPgSqlRepo *PokemonPgSqlRepo) GetById(pokemonId int) (repository.Pokemon, error) {
	pokemon := &repository.Pokemon{ID: pokemonId}

	err := pgDb.Model(pokemon).WherePK().Select()
	if err != nil {
		fmt.Println(err)
		return repository.Pokemon{}, errors.New("INTERNAL_SERVER_ERROR")
	}

	return *pokemon, nil
}

func (pokemonPgSqlRepo *PokemonPgSqlRepo) Insert(newPokemon repository.Pokemon) error {
	_, err := pgDb.Model(&newPokemon).Insert()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pokemonPgSqlRepo *PokemonPgSqlRepo) Update(updatedPokemon repository.Pokemon) error {
	_, err := pgDb.Model(&updatedPokemon).WherePK().Update()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pokemonPgSqlRepo *PokemonPgSqlRepo) Delete(pokemonId int) error {
	res, err := pgDb.Model(new (repository.Pokemon)).Where("id = ?", pokemonId).Delete()

	fmt.Println(res, err)

	if err != nil {
		return err
	}

	return nil
}
