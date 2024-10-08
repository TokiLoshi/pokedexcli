package main

import (
	"fmt"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide name of the pokemon you'd like to inspect")
	}
	pokemon := args[0]
	
	
	pokedex := cfg.caughtPokemon
	_, ok := pokedex[pokemon];
	if !ok {
		return fmt.Errorf("***** you have not caught %s ******", pokemon)
	}

	decoratedCommand, err := ascii.RenderTextOptions(pokemon, "red", "magenta")
	if err != nil {
		return fmt.Errorf("error generating ascii art: %w", err)
	}
	
	fmt.Println(decoratedCommand)
	fmt.Println("=========================================")
	name := pokedex[pokemon].Name
	height := pokedex[pokemon].Height
	weight := pokedex[pokemon].Weight
	stats := pokedex[pokemon].Stats
	allTypes := pokedex[pokemon].Types
	
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Weight: %d\n", weight)
	fmt.Printf("Stats: \n")
	for _, stat := range stats {
		// fmt.Printf("  -%v: %v\n", stat.Stat.Name, stats[value].BaseStat)
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, poketype := range allTypes {
		fmt.Printf("  -%v\n", poketype.Type.Name)
	}

	fmt.Println("=========================================")
	
	return nil
}