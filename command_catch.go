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
	fmt.Printf("let's try catch: %v\n", pokemon)

	pokemonCatch, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("issue accessing api: %w", err)
	}

experience := pokemonCatch.BaseExperience
fmt.Println(experience)

randomNumber := rand.Intn(10) * experience / 10

userGeneratedNumber := rand.Intn(100)

fmt.Printf("Pokemon rolled a: %v and user rolled a: %v\n", randomNumber, userGeneratedNumber)
if randomNumber > userGeneratedNumber {
	fmt.Printf("%v escaped!\n", pokemon)
} else {
	fmt.Printf("congratulations you caught: %v\n", pokemon)
}

	

	return nil
}