package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the anmes of 20 location areas in the pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "display the names of the previous 20 location areas in the pokemon world",
			callback:    commandMapb,
		},
	}
}
