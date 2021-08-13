package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pokemon-go/helpers"
	"pokemon-go/repository"
	"pokemon-go/repository/database"
	pokemonServices "pokemon-go/services/pokemon"
)

var (
	httpCode     int
	responseData helpers.Response
)

//// JSON REPOSITORY
//var pokemonRepository = &repository.PokemonJsonRepo{}

//PGSQL REPOSITORY
var pokemonPgSqlRepository = &database.PokemonPgSqlRepo{}

func GetAllPokemon(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		getAllPokemonService := pokemonServices.NewGetAllPokemonService(
			pokemonPgSqlRepository,
		)
		httpCode, responseData = getAllPokemonService.Run()
	} else {
		httpCode, responseData = helpers.PageNotFoundResponse()
	}

	helpers.SendResponse(resp, httpCode, responseData)
}

func GetPokemon(resp http.ResponseWriter, req *http.Request) {
	pokemonId, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err == nil {
		getPokemonService := pokemonServices.NewGetPokemonService(
			pokemonPgSqlRepository,
			pokemonId,
		)
		httpCode, responseData = getPokemonService.Run()
	} else {
		httpCode, responseData = helpers.InternalServerErrorResponse()
	}

	helpers.SendResponse(resp, httpCode, responseData)
}

func InsertNewPokemon(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var newPokemon repository.Pokemon
		json.NewDecoder(req.Body).Decode(&newPokemon)

		if newPokemon.Name != "" && len(newPokemon.Types) > 0 &&
			len(newPokemon.Weaknesses) > 0 {
			insertPokemonService := pokemonServices.NewInsertPokemonService(
				pokemonPgSqlRepository, newPokemon,
			)
			httpCode, responseData = insertPokemonService.Run()
		} else {
			httpCode, responseData = helpers.InvalidRequestResponse()
		}
	default:
		httpCode, responseData = helpers.MethodNotAllowedResponse()
	}

	helpers.SendResponse(resp, httpCode, responseData)
}

func UpdatePokemon(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var updatedPokemon repository.Pokemon
		json.NewDecoder(req.Body).Decode(&updatedPokemon)

		if updatedPokemon.ID > 0 && updatedPokemon.Name != "" &&
			len(updatedPokemon.Types) > 0 && len(updatedPokemon.Weaknesses) > 0 {
			updatePokemonService := pokemonServices.NewUpdatePokemonService(
				pokemonPgSqlRepository,
				updatedPokemon,
			)
			httpCode, responseData = updatePokemonService.Run()
		} else {
			httpCode, responseData = helpers.InvalidRequestResponse()
		}
	default:
		httpCode, responseData = helpers.MethodNotAllowedResponse()
	}

	helpers.SendResponse(resp, httpCode, responseData)
}

func DeletePokemon(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		pokemonId, err := strconv.Atoi(req.URL.Query().Get("id"))

		if err == nil {
			deletePokemonService := pokemonServices.NewDeletePokemonService(
				pokemonPgSqlRepository,
				pokemonId,
			)
			httpCode, responseData = deletePokemonService.Run()
		} else {
			httpCode, responseData = helpers.InternalServerErrorResponse()
		}
	default:
		httpCode, responseData = helpers.MethodNotAllowedResponse()
	}

	helpers.SendResponse(resp, httpCode, responseData)
}
