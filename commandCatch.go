package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, parameter []string) error {
	if len(parameter) == 0 {
		return fmt.Errorf("Please enter the pokemon to catch")
	}
	pokiUrl := "/pokemon/" + parameter[0]
	pokemon, err := cfg.pokiClient.GetPokemon(pokiUrl)
	if err != nil {
		return err
	}
	// fmt.Println(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if cfg.pokiuser.Experince > pokemon.BaseExperience {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		cfg.pokiuser.Pokemons[pokemon.Name] = pokemon
		cfg.pokiuser.Experince -= pokemon.BaseExperience
		fmt.Println(cfg.pokiuser.Pokemons)
		return nil
	}
	fmt.Printf("%v escaped!\n", pokemon.Name)
	expGained := rand.Intn(200)
	cfg.pokiuser.Experince += expGained
	fmt.Println(expGained)
	return nil
}
