package main

import (
	"fmt"
	"net/http"
	Controllers "pokemon-go/controllers"
)

func main() {
	routes()
	serves()
}

func serves() {
	fmt.Println("Server is connected to http://localhost:10000")
	http.ListenAndServe(":10000", nil)
}

func routes() {
	http.HandleFunc("/", Controllers.GetAllPokemon)
	http.HandleFunc("/pokemon", Controllers.GetPokemon)
	http.HandleFunc("/insert-pokemon", Controllers.InsertNewPokemon)
	http.HandleFunc("/update-pokemon", Controllers.UpdatePokemon)
	http.HandleFunc("/delete-pokemon", Controllers.DeletePokemon)
}
