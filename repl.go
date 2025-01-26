package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommands struct {
	name       string
	descrption string
	callback   func() error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		input = strings.ToLower(strings.Fields(strings.Trim(input, " "))[0])
		comm, d := getCommands()[input]
		if d {
			comm.callback()
		} else {
			fmt.Println("Unkown Command")
		}

	}
}

func cleanInputText(text string) []string {
	text = strings.ToLower(text)
	temp := strings.Split(text, " ")
	res := []string{}
	for _, str := range temp {
		if len(str) != 0 {
			res = append(res, str)
		}
	}
	return res
}

func getCommands() map[string]cliCommands {
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
	}
}
