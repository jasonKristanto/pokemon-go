package controllers

import (
	"encoding/json"
	"net/http"

	Pokemon "pokemon-go/data"
	Helpers "pokemon-go/helpers"
	PokemonServices "pokemon-go/services/pokemon"
)

func GetAllPokemon(resp http.ResponseWriter, req *http.Request) {
	var httpCode int
	var responseData Helpers.Response

	if req.URL.Path != "/" {
		httpCode, responseData = Helpers.PageNotFoundResponse()
	} else {
		httpCode, responseData = PokemonServices.GetAllPokemon()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func GetPokemon(resp http.ResponseWriter, req *http.Request)  {
	var httpCode int
	var responseData Helpers.Response

	httpCode, responseData = PokemonServices.GetPokemon(req)

	Helpers.SendResponse(resp, httpCode, responseData)
}

func InsertNewPokemon(resp http.ResponseWriter, req *http.Request)  {
	var httpCode int
	var responseData Helpers.Response

	switch req.Method {
	case "POST":
		var newPokemon Pokemon.Type
		json.NewDecoder(req.Body).Decode(&newPokemon)

		if newPokemon.Name == "" || len(newPokemon.Types) <= 0 || len(newPokemon.Weaknesses) <= 0 {
			httpCode, responseData = Helpers.InvalidRequestResponse()
		} else {
			httpCode, responseData = PokemonServices.InsertNewPokemon(newPokemon)
		}
	default:
		httpCode, responseData = Helpers.MethodNotAllowedResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func UpdatePokemon(resp http.ResponseWriter, req *http.Request)  {
	var httpCode int
	var responseData Helpers.Response

	switch req.Method {
	case "POST":
		var updatedPokemon Pokemon.Type
		json.NewDecoder(req.Body).Decode(&updatedPokemon)

		if updatedPokemon.ID == 0 || updatedPokemon.Name == "" || len(updatedPokemon.Types) <= 0 || len(updatedPokemon.Weaknesses) <= 0 {
			httpCode, responseData = Helpers.InvalidRequestResponse()
		} else {
			httpCode, responseData = PokemonServices.UpdatePokemon(updatedPokemon)
		}
	default:
		httpCode, responseData = Helpers.MethodNotAllowedResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}

func DeletePokemon(resp http.ResponseWriter, req *http.Request)  {
	var httpCode int
	var responseData Helpers.Response

	switch req.Method {
	case "DELETE":
		httpCode, responseData = PokemonServices.DeletePokemon(req)
	default:
		httpCode, responseData = Helpers.MethodNotAllowedResponse()
	}

	Helpers.SendResponse(resp, httpCode, responseData)
}
