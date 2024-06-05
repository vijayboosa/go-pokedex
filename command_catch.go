package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func getRandomNumber(baseExperience int) int {
	if baseExperience < 30 {
		baseExperience = 30
	} else if baseExperience > 350 {
		baseExperience = 350
	}

	var probabilityFactor float64

	if baseExperience <= 150 {
		probabilityFactor = 1 - (float64(baseExperience-30) / 150.0) // Probability decreases from 1 to 0.2 as baseExperience increases from 30 to 150
	} else {
		probabilityFactor = 0.2 / (float64(baseExperience-150)/200.0 + 1) // Probability decreases further as baseExperience increases from 150 to 350
	}

	if rand.Float64() < probabilityFactor {
		return 1
	}
	return 0
}

func commandCatch(cnf *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("Missing pokemon name to catch")
	}

	pokeName := args[0]

	pokemon, err := cnf.HttClient.CatchPokemon(pokeName)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokeName)

	if getRandomNumber(pokemon.BaseExperience) == 0 {
		fmt.Println(pokeName + " escaped!")
		return nil
	}

	cnf.PokeStorage[pokeName] = pokemon
	fmt.Println(pokeName + " was caught!")
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}
