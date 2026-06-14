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
	scanner := bufio.NewScanner(os.Stdin)
	var command string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		//skip if the input is empty
		if len(words) == 0 {
			continue
		}
		command = words[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}
