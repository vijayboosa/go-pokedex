package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vijayboosa/pokedex-cli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	Next        string
	Previous    string
	HttClient   pokeapi.Client
	PokeStorage map[string]pokeapi.Pokemon
}

func startRepl(cnf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		words := cleanText(scanner.Text())
		if len(words) == 0 {
			continue
		}

		wordCommand := words[0]
		extraArg := ""

		if len(words) > 1 {
			extraArg = words[1]
		}

		if command, ok := getCommand()[wordCommand]; ok {
			command.callback(cnf, extraArg)
		}

		fmt.Printf("\n")
	}
}

func cleanText(text string) []string {
	stringLower := strings.ToLower(text)

	words := strings.Fields(stringLower)
	return words
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": cliCommand{
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the",
			callback:    commandExit,
		},
		"map": cliCommand{
			name:        "map",
			description: "Get 20 pokemon location area",
			callback:    commandMap,
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Get previous 20 pokemon location area",
			callback:    commandMapb,
		},
		"explore": cliCommand{
			name:        "explore",
			description: "Get pokemons in specific area",
			callback:    commandExplore,
		},
		"catch": cliCommand{
			name:        "catch",
			description: "Catch pokemon by name",
			callback:    commandCatch,
		},
		"inspect": cliCommand{
			name:        "inspect",
			description: "Inspect the pokemon you caught",
			callback:    commandInspect,
		},
		"pokedex": cliCommand{
			name:        "pokedex",
			description: "See all the pokemon you caugth",
			callback:    commandPokedex,
		},
	}
}
