package main

import "fmt"

func commandMapf(cfg *config, parameter []string) error {
	// locationRes , ok :=
	locRes, err := cfg.pokiClient.ListLocations(cfg.nextUrl)
	if err != nil {
		return err
	}

	cfg.nextUrl = locRes.Next
	cfg.previousUrl = locRes.Previous
	fmt.Println(string(*cfg.nextUrl))
	for _, loc := range locRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, parameter []string) error {
	locRes, err := cfg.pokiClient.ListLocations(cfg.previousUrl)
	if err != nil {
		return err
	}

	cfg.nextUrl = locRes.Next
	cfg.previousUrl = locRes.Previous
	for _, loc := range locRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
