package services

import (
	"pokemon-go/helpers"
)

type IBasePokemonService interface {
	Run() (int, helpers.Response)
}
