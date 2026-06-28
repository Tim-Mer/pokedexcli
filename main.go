package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Tim-Mer/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cli := getCommands()
	config := Config{
		CurrentPage: 1,
		runningCache: pokecache.NewCache(time.Second * 60),
		catchHelper: 0,
	}
	configPtr := &config

	for {
		//fmt.Printf("config: %v\n", config)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", scanner.Err())
			return
		}
		input := cleanInput(scanner.Text())
		command := cli[input[0]].callback
		if len(input) > 1 {
			//add arguments so they can be passed to commands that require them
			config.arguments = []byte(input[1])
		}

		if command == nil {
			fmt.Println("Unknown command")
		} else {
			command(configPtr)
		}

	}
}
