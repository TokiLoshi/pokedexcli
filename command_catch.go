package main

import (
	"fmt"
	"math/rand"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandCatch(cfg *config, args ...string) error {
	

	if len(args) < 1 {
		return fmt.Errorf("please choose a pokemon")
	}
	pokemon := args[0]

	// decoratedCommand, err := ascii.RenderTextOptions("catch " + pokemon, "magenta", "red")
	// if err != nil {
	// 	return fmt.Errorf("error rndering ascii command: %w", err)
	// }
	// fmt.Println(decoratedCommand)

	pokemonCatch, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("issue accessing api: %w", err)
	}

	experience := pokemonCatch.BaseExperience

	randomNumber := rand.Intn(10) * experience / 10



	userGeneratedNumber := rand.Intn(100)
	fmt.Println("================================================")
	fmt.Printf("******* Throwing a pokeball at %v *******\n", pokemonCatch.Name)
	fmt.Println("================================================")
	// Evaluate if user caught the pokemon
	if randomNumber > userGeneratedNumber {
		sad, err := ascii.RenderTextOptions(pokemonCatch.Name + " escaped!", "red", "red")
		if err != nil {
			return fmt.Errorf("error generating ascii text: %w", err)
		}
		// fmt.Printf("%v escaped!\n", pokemonCatch.Name)
		fmt.Println(sad)
	} else {
		happy, err := ascii.RenderTextOptions(pokemonCatch.Name + " caught!", "green", "green")
		if err != nil {
			return fmt.Errorf("error generating ascii text: %w", err)
		}
		fmt.Println(happy)
		fmt.Printf("You may now inspect %v with the inspect command\n", pokemonCatch.Name)
	}

	cfg.caughtPokemon[pokemonCatch.Name] = pokemonCatch

	return nil
}