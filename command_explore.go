package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if cfg == nil {
		return errors.New("config argument can't be nil")
	}
	if len(args) < 1 {
		return errors.New("explore needs an argument")
	}
	name := args[0]
	location, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	if len(location.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, encounter := range location.PokemonEncounters {
			fmt.Printf("  - %s\n", encounter.Pokemon.Name)
		}
	} else {
		fmt.Println("Not found pokemons in area")
	}

	return nil
}
