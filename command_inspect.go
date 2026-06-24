package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("inspect command needs the pokemon name as an argument")
	}
	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Height)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")

	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
