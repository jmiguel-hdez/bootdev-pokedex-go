package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config) error {
	if cfg == nil {
		return errors.New("config can't be nil")
	}

	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
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
