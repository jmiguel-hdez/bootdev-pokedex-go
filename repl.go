package main

import (
	"bufio"
	"fmt"
	"github.com/jmiguel-hdez/bootdev-pokedex-go/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func cleanInput(text string) []string {
	lowercase_text := strings.ToLower(text)
	return strings.Fields(lowercase_text)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		//skip if the input is empty
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
