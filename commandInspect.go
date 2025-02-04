package main

import "fmt"

func commandInspect(cfg *config, parameter []string) error {
	if len(parameter) == 0 {
		return fmt.Errorf("Enter an pokemon name")
	}
	pokeTemp, ok := cfg.pokiuser.Pokemons[parameter[0]]
	if !ok {
		return fmt.Errorf("You haven't caught this pokemon!!")
	}
	fmt.Printf("Name : %v\n", pokeTemp.Name)
	fmt.Printf("Height : %v\n", pokeTemp.Height)
	fmt.Printf("Weight : %v\n", pokeTemp.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokeTemp.Stats {
		fmt.Printf("\t-%v : %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokeTemp.Types {
		fmt.Printf("\t- %v \n", t.Type.Name)
	}
	return nil
}
