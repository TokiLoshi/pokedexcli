package main

import (
	"time"

	"github.com/TokiLoshi/pokedexcli/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(8 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)

	
}


