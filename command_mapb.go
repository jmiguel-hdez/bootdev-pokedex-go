package main

import (
	"errors"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-pokedex-go/internal/pokedex"
)

func commandMapb(config *Config) error {
	if config == nil {
		return errors.New("config can't be nil")
	}
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := pokedex.Mapb(config.Previous)
	if err != nil {
		return err
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	config.Next = locations.Next
	config.Previous = locations.Previous
	return nil
}
