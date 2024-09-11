package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please choose a pokemon")
	}
	pokemon := args[0]

	pokemonCatch, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("issue accessing api: %w", err)
	}

	experience := pokemonCatch.BaseExperience

	randomNumber := rand.Intn(10) * experience / 10

	userGeneratedNumber := rand.Intn(100)
	
	// Evaluate if user caught the pokemon
	if randomNumber > userGeneratedNumber {
		fmt.Printf("%v escaped!\n", pokemonCatch.Name)
	} else {
		fmt.Printf("congratulations you caught: %v\n", pokemonCatch.Name)
	}

	cfg.caughtPokemon[pokemonCatch.Name] = pokemonCatch

	return nil
}