package main

import (
	"fmt"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandHelp(cfg *config, args ...string) error {

	commands := getCommands()

	decoratedCommand, err := ascii.RenderTextOptions("menu", "green", "green")
	if err != nil {
		return fmt.Errorf("error decorating help command: %w", err)
	}
	fmt.Println(decoratedCommand)
	for _, command := range commands {
		fmt.Printf("%v:\n", command.name)
		fmt.Printf("   - %v\n", command.description)
	
	}
	return nil	
}