package main

import "fmt"

func commandPokedex(cnf *config, args ...string) error {

	fmt.Println("Your Pokedex:")
	for key, _ := range cnf.PokeStorage {
		fmt.Println(" - " + key)
	}

	return nil
}
