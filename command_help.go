package main

import "fmt"

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for key, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", key, cmd.description)
	}
	return nil
}
