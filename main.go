package main

import (
	"bufio"
	"fmt"
	"os"
    "github.com/primawk/pokedexgo/repl"
)

func main() {
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)

	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for scanner.Scan() {
		msg := fmt.Sprintf("Your command was: %s", repl.CleanInput(scanner.Text()))
		fmt.Println(msg)
		fmt.Print("Pokedex > ")
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}