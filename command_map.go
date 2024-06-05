package main

import (
	"fmt"
)

func commandMap(cnf *config, args ...string) (err error) {

	if cnf.Next == "" {
		cnf.Next = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
	}

	result, err := cnf.HttClient.FetchLocation(cnf.Next)

	if err != nil {
		return
	}

	cnf.Next = result.Next
	cnf.Previous = result.Previous
	for _, location := range result.Results {
		fmt.Println(location.Area)
	}

	return nil
}

func commandMapb(cnf *config, args ...string) (err error) {

	if cnf.Previous == "" {
		cnf.Next = ""
		fmt.Print("no pokemons to show")
		return
	}

	result, err := cnf.HttClient.FetchLocation(cnf.Previous)

	if err != nil {
		fmt.Println("Failed get the result")
		return
	}

	cnf.Next = result.Next
	cnf.Previous = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Area)
	}

	return nil
}
