package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	fmt.Println("getting 20 locations")
	locales, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("error with locations: %w", err)
	}
	
	// Update last location 
	cfg.nextLocationsURL = locales.Next
	cfg.previousLocationsURL = locales.Previous

	for _, location := range locales.Results {
		fmt.Println(location.Name)
	}
	
	return nil

}

func commandMapb(cfg *config) error {

	fmt.Println("going back 20 locations")

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