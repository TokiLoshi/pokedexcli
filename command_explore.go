package main

import (
	"fmt"
	"strconv"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) < 1 {
		return fmt.Errorf("please provide a location")
	}

	locationName := args[0]
	
	decoratedCommand, err := ascii.RenderTextOptions("exploring", "yellow", "yellow")
	if err != nil {
		return fmt.Errorf("error rendering ascii art: %w", err)
	}
	fmt.Println(decoratedCommand)

	
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return fmt.Errorf("error searching for pokemon: %w", err)
	}

	

	for _, encounter := range location.PokemonEncounters{
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	
	amount := strconv.Itoa(len(location.PokemonEncounters))

	findings, err := ascii.RenderTextOptions(amount + " found", "green", "green")
	if err != nil {
		return fmt.Errorf("error rendering ascii art: %w", err)
	}
	fmt.Println(findings)
	return nil
}