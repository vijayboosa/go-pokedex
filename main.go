package main

import (
	"github.com/vijayboosa/pokedex-cli/pokeapi"
)

func main() {
	cnf := config{
		HttClient:   pokeapi.NewClient(),
		PokeStorage: map[string]pokeapi.Pokemon{},
	}
	startRepl(&cnf)
}
