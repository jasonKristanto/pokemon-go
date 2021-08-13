package services

import (
	"github.com/stretchr/testify/mock"

	mock2 "pokemon-go/tests/mock"
)

var pokemonRepositoryMock = &mock2.PokemonJsonMockRepo{Mock: mock.Mock{}}
