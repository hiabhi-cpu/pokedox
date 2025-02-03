package main

import "fmt"

func commandExplore(con *config, parameter []string) error {
	if len(parameter) == 0 {
		return fmt.Errorf("Please enter the location name")
	}
	locUrl := "/location-area/" + parameter[0]
	listExpRes, err := con.pokiClient.ListLocationExplore(locUrl)
	if err != nil {
		return err
	}
	// fmt.Println(len(listExpRes.PokemonEncounters))
	for _, v := range listExpRes.PokemonEncounters {
		fmt.Println(" - " + v.Pokemon.Name)
	}
	return nil
}
