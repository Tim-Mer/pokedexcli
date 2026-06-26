package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Tim-Mer/pokedexcli/internal/pokeapi"
	"github.com/Tim-Mer/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	CurrentPage  int
	runningCache *pokecache.Cache
	arguments    []byte
}

var urlLocationAPI = "https://pokeapi.co/api/v2/location-area/"
var urlPokemonAPI = "https://pokeapi.co/api/v2/pokemon/"

func getCommands() map[string]cliCommand {
	list := map[string]cliCommand{
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
			name:        "map",
			description: "Show the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back a page on the map",
			callback:    backPage,
		},
		"explore": {
			name:        "explore",
			description: "Explore the current area and display the current available pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
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

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	list := getCommands()
	for command := range list {
		fmt.Printf("%v: %v\n", list[command].name, list[command].description)
	}
	return nil
}

func commandMap(config *Config) error {
	var err error

	//Getting all the locations from the api/cache
	for i := 1; i <= 20; i++ {
		pageNum := ((config.CurrentPage - 1) * 20) + i
		tmpurl := urlLocationAPI + strconv.Itoa(pageNum) + "/"
		var res []byte
		var found bool

		//chech cache first for location data
		if res, found = config.runningCache.Get(tmpurl); !found {
			//if not found get it from the api
			res, err = pokeapi.GetLocation(tmpurl)
			if err != nil {
				return err
			}
			// and add it to the cache
			config.runningCache.Add(tmpurl, res)
		}
		//Print location
		fmt.Printf("%s\n", string(res))
	}

	//fmt.Printf("Current page: %v\n", config.CurrentPage)
	config.CurrentPage++
	return err
}

func backPage(config *Config) error {
	if config.CurrentPage <= 2 {
		fmt.Print("you're on the first page\n")
		return nil
	} else {
		config.CurrentPage -= 2
	}
	return commandMap(config)
}

func commandExplore(config *Config) error {
	tmpurl := urlLocationAPI + string(config.arguments) + "/"
	var res []byte
	var found bool
	var err error

	//chech cache first for explore data
	if res, found = config.runningCache.Get(tmpurl); !found {
		//if not found get it from the api
		res, err = pokeapi.ExploreLocation(tmpurl)
		if err != nil {
			return err
		}
		// and add it to the cache
		config.runningCache.Add(tmpurl, res)
	}
	//need to print out the result
	for _, name := range res {
		fmt.Printf("%s", string(name))
	}
	return nil
}

func commandCatch(config *Config) error {
	tmpurl := urlPokemonAPI + string(config.arguments) + "/"
	return nil
}
