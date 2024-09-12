package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TokiLoshi/pokedexcli/internal/ascii"
	"github.com/TokiLoshi/pokedexcli/internal/pokeapi"
)


type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	previousLocationsURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	asciiArt, err := ascii.RenderTextOptions("Pokedex", "yellow", "red")
	if err != nil {
		fmt.Println("error loading ascii text: %w", err)
	}

	fmt.Println(asciiArt)
	fmt.Println("=========================================")
	fmt.Println("              Instructions:              ")
	fmt.Println("=========================================")
	fmt.Println("")	
	fmt.Println("1. Move through the map 20 locations at a time with -'mapf'")
	fmt.Println("")
	fmt.Println("2. Move back through the map with -'mapb'")
	fmt.Println("")
	fmt.Println("3. Explore an area with -'explore <area_name>'")
	fmt.Println("")
	fmt.Println("4. If you find pokemon try catch one with -'catch <pokemon_name>'")
	fmt.Println("")
	fmt.Println("5. If you are successful you can inspect them -'inspect <pokemon_name>'")
	fmt.Println("")
	fmt.Println("6. Admire your pokedex -'pokedex'")
	fmt.Println("")
	fmt.Println("View the menu at any time with -'help'")
	fmt.Println("")
	fmt.Println("=========================================")


	for {
		fmt.Print("Pokedex> ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		// Adjusted logic to handle multiple args
		if len(words) > 1 {
			args = words[1:]
		}

		if command, ok := getCommands()[commandName]; ok {
			// at this point we get "help" or exit 
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
	} else {
		fmt.Println("=========================================")
		fmt.Println("Your options are:")
		fmt.Println("=========================================")
		fmt.Println(" - help\n - exit\n - mapf\n - mapb\n - explore <location-name>\n - catch <pokemon-name>\n - inspect <pokemon-name>\n - pokedex")
		fmt.Println("=========================================")
		continue
	}
}
}

func cleanInput(text string) []string {
	commandInput := strings.ToLower(text)
	words := strings.Fields(commandInput)
	return words
}

func getCommands() map[string] cliCommand{
	return map[string]cliCommand {
		"catch": {
			name: "catch <pokemon_name>",
			description: "Tries to catch a pokemon",
			callback: commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Finds Pokemon in the area",
			callback: commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name: "inspect",
			description: "Inspects a caught pokemon",
			callback: commandInspect,
		},
		"mapf": {
			name: "mapf",
			description: "Displays the next 20 location areas in Pokemon world",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback: commandMapb,
		},
		"pokedex": {
			name: "pokedex",
			description: "Prints out the contents of your pokedex",
			callback: commandPokedex,
		},

	}
}