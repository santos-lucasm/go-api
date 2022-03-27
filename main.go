package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Name string `json:"name"`
}

func main() {

	response := httpRequest("http://pokeapi.co/api/v2/pokedex/kanto/")
	responseData := processResponse(response)

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(string(responseObject.Name))
}
