package main

import (
	"fmt"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandMapf(cfg *config, args ...string) error {
	decoratedCommand, err := ascii.RenderTextOptions("forward", "magenta", "magenta")
	if err != nil {
		return fmt.Errorf("error generating ascii art")
	}
	fmt.Println(decoratedCommand)

	locales, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("error with locations: %w", err)
	}
	
	cfg.nextLocationsURL = locales.Next
	cfg.previousLocationsURL = locales.Previous

	for _, location := range locales.Results {
		fmt.Println(location.Name)
	}
	
	return nil

}

func commandMapb(cfg *config, args ...string) error {

	decoratedCommand, err := ascii.RenderTextOptions("back", "cyan", "cyan")
	if err != nil {
		return fmt.Errorf("error generating ascii art")
	}
	fmt.Println(decoratedCommand)

	locales, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return fmt.Errorf("error going back locations: %w", err)
	}

	cfg.nextLocationsURL = locales.Next
	cfg.previousLocationsURL = locales.Previous

	for _, location := range locales.Results {
		fmt.Println(location.Name)
	}

	return nil
}