package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	Pokemon "pokemon-go/data"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []Pokemon.Type `json:"data"`
}

func getAllPokemon(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	if req.URL.Path != "/" {
		resp.WriteHeader(404)

		jsonInBytes, err := json.Marshal(Response {
			Status: "error",
			Message: "PAGE_NOT_FOUND",
		})

		if err != nil {
			resp.WriteHeader(500)

			jsonInBytes, err = json.Marshal(Response {
				Status: "error",
				Message: "INTERNAL_SERVER_ERROR",
			})
		}

		resp.Write(jsonInBytes)
		return
	}

	jsonInBytes, err := json.Marshal(Response {
		Status: "success",
		Message: "DATA_FOUND",
		Data: *(Pokemon.GetData()),
	})

	if err != nil {
		resp.WriteHeader(500)

		jsonInBytes, err = json.Marshal(Response {
			Status: "error",
			Message: "INTERNAL_SERVER_ERROR",
		})
	}

	resp.Write(jsonInBytes)
}

func getPokemon(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	var data = Response{}
	pokemons := *(Pokemon.GetData())
	pokemonName := req.URL.Query().Get("name")

	if len(pokemonName) > 0 {
		for _, value := range pokemons {
			if strings.ToLower(value.Name) == strings.ToLower(pokemonName) {
				data = Response {
					Status: "success",
					Message: "DATA_FOUND",
					Data: []Pokemon.Type{value},
				}
				break
			}
		}
	}

	jsonInBytes, err := json.Marshal(data)

	if data.Status != "success" {
		resp.WriteHeader(404)

		jsonInBytes, err = json.Marshal(Response {
			Status: "error",
			Message: "DATA_NOT_FOUND",
		})
	}

	if err != nil {
		resp.WriteHeader(500)

		jsonInBytes, err = json.Marshal(Response {
			Status: "error",
			Message: "INTERNAL_SERVER_ERROR",
		})
	}

	resp.Write(jsonInBytes)
}

func insertNewPokemon(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		var newPokemon Pokemon.Type
		json.NewDecoder(req.Body).Decode(&newPokemon)

		if newPokemon.Name == "" || len(newPokemon.Types) <= 0 || len(newPokemon.Weaknesses) <= 0 {
			resp.WriteHeader(400)

			jsonInBytes, _ := json.Marshal(Response {
				Status: "error",
				Message: "INVALID_REQUEST",
			})

			resp.Write(jsonInBytes)
			return
		}

		pokemons := Pokemon.GetData()
		lastPokemon := (*pokemons)[len(*pokemons) - 1]
		newPokemon.ID = lastPokemon.ID + 1

		*pokemons = append(*pokemons, newPokemon)

		jsonInBytes, err := json.Marshal(Response {
			Status: "success",
			Message: "INSERT_SUCCESSFUL",
		})

		if err != nil {
			resp.WriteHeader(500)

			jsonInBytes, _ = json.Marshal(Response {
				Status: "error",
				Message: "INTERNAL_SERVER_ERROR",
			})
		}

		resp.Write(jsonInBytes)
	default:
		resp.WriteHeader(405)

		jsonInBytes, _ := json.Marshal(Response {
			Status: "error",
			Message: "METHOD_NOT_ALLOWED",
		})

		resp.Write(jsonInBytes)
	}
}

func updatePokemon(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		isPokemonExisted := false
		var existedPokemon Pokemon.Type
		json.NewDecoder(req.Body).Decode(&existedPokemon)

		if existedPokemon.Name == "" || len(existedPokemon.Types) <= 0 || len(existedPokemon.Weaknesses) <= 0 {
			resp.WriteHeader(400)

			jsonInBytes, _ := json.Marshal(Response {
				Status: "error",
				Message: "INVALID_REQUEST",
			})

			resp.Write(jsonInBytes)
			return
		}

		jsonInBytes, _ := json.Marshal(Response {
			Status: "error",
			Message: "DATA_NOT_FOUND",
		})

		pokemons := Pokemon.GetData()
		for index, value := range *pokemons {
			if strings.ToLower(value.Name) == strings.ToLower(existedPokemon.Name) {
				isPokemonExisted = true
				(*pokemons)[index] = Pokemon.Type {
					ID: value.ID,
					Name: existedPokemon.Name,
					Types: existedPokemon.Types,
					Weaknesses: existedPokemon.Weaknesses,
				}
				break
			}
		}

		if isPokemonExisted {
			jsonInBytes, _ = json.Marshal(Response {
				Status: "success",
				Message: "UPDATE_SUCCESSFUL",
			})
		}

		resp.Write(jsonInBytes)
	default:
		resp.WriteHeader(405)

		jsonInBytes, _ := json.Marshal(Response {
			Status: "error",
			Message: "METHOD_NOT_ALLOWED",
		})

		resp.Write(jsonInBytes)
	}
}

func deletePokemon(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "DELETE":
		isPokemonExisted := false
		pokemonId, _ := strconv.Atoi(req.URL.Query().Get("id"))

		jsonInBytes, _ := json.Marshal(Response {
			Status: "error",
			Message: "DATA_NOT_FOUND",
		})

		pokemons := Pokemon.GetData()
		for index, value := range *pokemons {
			if value.ID == pokemonId {
				isPokemonExisted = true
				copy((*pokemons)[index:], (*pokemons)[index+1:])
				(*pokemons)[len(*pokemons)-1] = Pokemon.Type{}
				*pokemons = (*pokemons)[:len(*pokemons)-1]
				break
			}
		}

		if isPokemonExisted {
			jsonInBytes, _ = json.Marshal(Response {
				Status: "success",
				Message: "DELETE_SUCCESSFUL",
			})
		}

		resp.Write(jsonInBytes)
	default:
		resp.WriteHeader(405)

		jsonInBytes, _ := json.Marshal(Response {
			Status: "error",
			Message: "METHOD_NOT_ALLOWED",
		})

		resp.Write(jsonInBytes)
	}
}

func routes() {
	http.HandleFunc("/", getAllPokemon)
	http.HandleFunc("/pokemon", getPokemon)
	http.HandleFunc("/insert-pokemon", insertNewPokemon)
	http.HandleFunc("/update-pokemon", updatePokemon)
	http.HandleFunc("/delete-pokemon", deletePokemon)
}

func handleRequests() {
	routes()

	fmt.Println("Server is connected to http://localhost:10000")
	http.ListenAndServe(":10000", nil)
}

func main() {
	handleRequests()
}
