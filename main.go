package main

import (
	"time"

	"github.com/TokiLoshi/pokedexcli/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(8 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)

	
}


