package main

import (
	"fmt"
	"os"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
)

func commandExit(cfg *config, args ...string) error {
	decoratedCommand, err := ascii.RenderTextOptions("Byeee", "cyan", "cyan")
	if err != nil {
		return fmt.Errorf("error generating ascii text")
	}
	fmt.Println(decoratedCommand)
	os.Exit(0)
	return nil
}