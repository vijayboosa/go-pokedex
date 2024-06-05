package main

import "fmt"

func commandHelp(cnf *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, value := range getCommand() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
