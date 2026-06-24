package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore <location_name>",
			description: "display location area information including the possible pokemon encounters, needs <location_name> as parameter",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "receives a name of a Pokemon as an argument. It attempts to catch the pokemon, if the catch succeds the pokemon is added to the pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "inspect pokemon information(pokemon needs to be on pokedex)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "prints a list of all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}
