package main

import "fmt"

func commandpokedox(cfg *config, parameter []string) error {
	if len(cfg.pokiuser.Pokemons) == 0 {
		return fmt.Errorf("No Pokemon is caught by you")
	}
	fmt.Println("Your Pokedox:")
	for _, v := range cfg.pokiuser.Pokemons {
		fmt.Println("\t - ", v.Name)
	}
	return nil
}
