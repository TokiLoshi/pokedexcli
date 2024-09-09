package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func commandHelp() {
	fmt.Println("ellllllppp")
	
}

func commandExit() {
	fmt.Println("exit!")
	os.Exit(0)
}

func commandMap() {
	fmt.Println("getting 20 locations")
}

func commandMapb() {
	fmt.Println("going back 20 locations")
}

func configure() map[string] cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name: "map",
			description: "Displays names of 20 location areas in Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Takes you back a command",
			callback: commandMapb,
		},
	}
}