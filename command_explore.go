package main

import (
	"fmt"
)

// func (c *Client) ListPokemon(pageURL *string, area string) error) {
func commandExplore(cfg *config, args ...string) error {
	// If less than 1 args error - we need a location 
	if len(args) < 1 {
		return fmt.Errorf("please provide a location")
	}
	// set location name 
	locationName := args[0]
	
	fmt.Printf("Great, let's explore: %v", locationName)

	// make a call to the api enpoint 
	// handle the error 
	// iterate over the areas and read it out
	
	
	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return fmt.Errorf("error searching for pokemon: %w", err)
	}
	fmt.Println("Found pokemon...")
	for _, encounter := range location.PokemonEncounters{
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	return nil
}