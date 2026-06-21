package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowercase_text := strings.ToLower(text)
	return strings.Fields(lowercase_text)
}

func startRepl() {
	var config Config
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
			err := command.callback(&config)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing command:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
