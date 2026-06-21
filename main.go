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

		if command == nil {
			fmt.Println("Unknown command")
		} else {
			command(configPtr)
		}

	}
}
