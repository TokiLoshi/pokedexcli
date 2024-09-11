package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TokiLoshi/pokedexcli/internal/pokeapi"
)


type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	previousLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex Comand Line Center")
	fmt.Println("******************")


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
		fmt.Println("I'm not sure what you mean... we seem to have a misunderstanding")
		fmt.Println("Your options are:\n -help\n -exit\n -mapf\n -mapb\n -explore <location-name>\n -catch <pokemon-name>")
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
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name: "map",
			description: "Displays names of 20 location areas in Pokemon world",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Takes you back a command",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Finds Pokemon in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Tries to catch a pokemon",
			callback: commandCatch,
		},
	}
}