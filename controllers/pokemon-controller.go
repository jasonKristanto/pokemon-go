package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
	PokemonServices "pokemon-go/services/pokemon"
)

var (
	httpCode     int
	responseData Helpers.Response
)

var pokemonService = PokemonServices.PokemonService{
	Pokemons: Pokemon.GetData(),
}

func GetAllPokemon(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		getAllPokemonService := PokemonServices.NewGetAllPokemonService(pokemonService)
		httpCode, responseData = getAllPokemonService.Run()
	} else {
		httpCode, responseData = Helpers.PageNotFoundResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func GetPokemon(resp http.ResponseWriter, req *http.Request) {
	pokemonId, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err == nil {
		getPokemonService := PokemonServices.NewGetPokemonService(pokemonService, pokemonId)
		httpCode, responseData = getPokemonService.Run()
	} else {
		httpCode, responseData = Helpers.InternalServerErrorResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func InsertNewPokemon(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var newPokemon Pokemon.Type
		json.NewDecoder(req.Body).Decode(&newPokemon)

		if newPokemon.Name != "" && len(newPokemon.Types) > 0 && len(newPokemon.Weaknesses) > 0 {
			insertPokemonService := PokemonServices.NewInsertPokemonService(pokemonService, newPokemon)
			httpCode, responseData = insertPokemonService.Run()
		} else {
			httpCode, responseData = Helpers.InvalidRequestResponse()
		}
	default:
		httpCode, responseData = Helpers.MethodNotAllowedResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func UpdatePokemon(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var updatedPokemon Pokemon.Type
		json.NewDecoder(req.Body).Decode(&updatedPokemon)

		if updatedPokemon.ID > 0 && updatedPokemon.Name != "" && len(updatedPokemon.Types) > 0 && len(updatedPokemon.Weaknesses) > 0 {
			updatePokemonService := PokemonServices.NewUpdatePokemonService(pokemonService, updatedPokemon)
			httpCode, responseData = updatePokemonService.Run()
		} else {
			httpCode, responseData = Helpers.InvalidRequestResponse()
		}
	default:
		httpCode, responseData = Helpers.MethodNotAllowedResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func DeletePokemon(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "DELETE":
		pokemonId, err := strconv.Atoi(req.URL.Query().Get("id"))

		if err == nil {
			deletePokemonService := PokemonServices.NewDeletePokemonService(pokemonService, pokemonId)
			httpCode, responseData = deletePokemonService.Run()
		} else {
			httpCode, responseData = Helpers.InternalServerErrorResponse()
		}
	default:
		httpCode, responseData = Helpers.MethodNotAllowedResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}
