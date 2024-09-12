package main

import (
	"fmt"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandPokedex(cfg *config, args ...string) error {

	pokedex := cfg.caughtPokemon 
	decoratedCommand, err := ascii.RenderTextOptions("Your Pokedex: ", "magenta", "magenta")
	if err != nil {
		return fmt.Errorf("error rendering ascii art: %w", err)
	}
	fmt.Println(decoratedCommand)
	fmt.Println("=========================================")
	if len(pokedex) == 0 {
		fmt.Printf("this is sad... your pokedex is empty,\ntry catch something.\nremember you gotta catch them all!\n")
	}

	for pokemon := range pokedex {
		fmt.Printf("- %v\n", pokemon)
	}
	fmt.Println("=========================================")
	return nil
}