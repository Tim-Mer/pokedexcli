package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cli := getCommands()
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", scanner.Err())
			return
		}
		input := cleanInput(scanner.Text())[0]
		command := cli[input].callback
		if command == nil {
			fmt.Println("Unknown command")
		} else {
			command()
		}
		
	}
}