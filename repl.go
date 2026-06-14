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
	var input string
	var command string
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			input = scanner.Text()
			words := cleanInput(input)
			command = words[0]
			fmt.Println("your command was: " + command)
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
