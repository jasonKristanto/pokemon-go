package main

import (
	"fmt"
	"net/http"
	"strconv"

	"pokemon-go/config"
	"pokemon-go/controllers"
)

func main() {
	config.Init()
	routes()
	serves()
}

func serves() {
	serverPort := config.Configuration.Server.Port
	serverHost := config.Configuration.Server.Host
	server := fmt.Sprintf("http://%s:%d", serverHost, serverPort)
	fmt.Println("Server is connected to ", server)
	http.ListenAndServe(":" + strconv.Itoa(serverPort), nil)
}

func routes() {
	http.HandleFunc("/", controllers.GetAllPokemon)
	http.HandleFunc("/pokemon", controllers.GetPokemon)
	http.HandleFunc("/insert-pokemon", controllers.InsertNewPokemon)
	http.HandleFunc("/update-pokemon", controllers.UpdatePokemon)
	http.HandleFunc("/delete-pokemon", controllers.DeletePokemon)
}
