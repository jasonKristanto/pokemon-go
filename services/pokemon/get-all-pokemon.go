package services

import (
	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
)

func GetAllPokemon() (int, Helpers.Response) {
	return Helpers.SuccessResponse("DATA_FOUND", *(Pokemon.GetData()))
}