package main

type cliCommands struct {
	name       string
	descrption string
	callback   func(*config) error
}

func getCommands(con *config) map[string]cliCommands {
	return map[string]cliCommands{
		"exit": {
			name:       "exit",
			descrption: "Exit the Pokedex",
			callback:   commandExit,
		},
		"help": {
			name:       "help",
			descrption: "Displays a help command",
			callback:   commandHelp,
		},
		"map": {
			name:       "map",
			descrption: "Show locations",
			callback:   commandMap,
		},
		"mapb": {
			name:       "mapb",
			descrption: "Show previous locations",
			callback:   commandMapBack,
		},
	}
}
