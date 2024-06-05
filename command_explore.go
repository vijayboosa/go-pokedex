package main

import (
	"errors"
	"fmt"
)

func commandExplore(cnf *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("Missing location name to explore")
	}

	area := args[0]

	data, err := cnf.HttClient.ExpoloreArea(area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}
	return nil
}
