package main

import (
	"errors"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-pokedex-go/internal/pokedex"
)

func commandMap(config *Config) error {
	if config == nil {
		return errors.New("config can't be nil")
	}
	locations, err := pokedex.Map(config.Next)
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
