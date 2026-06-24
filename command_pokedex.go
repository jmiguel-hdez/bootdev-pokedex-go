package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if cfg == nil {
		return errors.New("cfg can't be nil")
	}

	if len(cfg.pokedex) == 0 {
		return errors.New("pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
