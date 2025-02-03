package main

import (
	"time"

	"github.com/hiabhi-cpu/pokedox/internal/pokeapi"
)

func main() {
	pokeClienttemp := pokeapi.NewClient(time.Second*5, time.Minute*2)

	cfg := config{
		pokiClient: pokeClienttemp,
		pokiuser:   pokeapi.NewUser(),
	}

	startRepl(&cfg)
}
