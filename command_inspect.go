package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandInspect(cnf *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Missing pokemon name to inspect")
	}

	name := args[0]

	if data, ok := cnf.PokeStorage[strings.ToLower(name)]; ok {
		fmt.Println("Name: ", data.Name)
		fmt.Println("Height: ", data.Height)
		fmt.Println("Weight: ", data.Weight)
		fmt.Println("Stats:")
		for _, stat := range data.Stats {
			fmt.Printf(" -%s: %d", stat.Stat.Name, stat.BaseStat)
			fmt.Println()
		}
		fmt.Println("Types:")
		for _, t := range data.Types {
			fmt.Printf(" - %s", t.Type.Name)
			fmt.Println()
		}

		return nil
	}

	fmt.Println("you have not caught that pokemon")
	return nil
}
