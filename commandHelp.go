package main

import (
	"fmt"
)

func commandHelp(con *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands(con) {
		fmt.Printf("%s: %s\n", cmd.name, cmd.descrption)
	}
	fmt.Println()
	return nil
}
