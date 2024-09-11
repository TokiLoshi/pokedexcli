package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("User wants to see their pokedex")
	pokedex := cfg.caughtPokemon 
	fmt.Printf("Your Pokedex:\n")
	if len(pokedex) == 0 {
		fmt.Printf("this is sad... your pokedex is empty,\ntry catch something.\nremember you gotta catch them all!\n")
	}
	for pokemon := range pokedex {
		fmt.Printf("- %v\n", pokemon)
	}
	return nil
}