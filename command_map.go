package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	if cfg == nil {
		return errors.New("config can't be nil")
	}
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
