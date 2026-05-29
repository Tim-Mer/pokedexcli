package main

import (
	"os"
	"strings"
	"fmt"
)

type cliCommand struct {
	name		string
	description string
	callback	func() error
}

func getCommands() map[string]cliCommand {
	list := map[string]cliCommand{
		"help": {
			name:			"help",
			description: 	"Displays a help message",
			callback: 		commandHelp,
		},
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
	}
	return list
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{""}
	}
	text = strings.ToLower(text)
	output := strings.Fields(text)

	return output
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	list := getCommands()
	for command := range list {
		fmt.Printf("%v: %v\n", list[command].name, list[command].description)
	}
	return nil
}