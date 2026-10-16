package main

import (
	"bufio"
	"fmt"
	"github.com/primawk/pokedexgo/internal"
	"github.com/primawk/pokedexgo/internal/pokecache"
	"github.com/primawk/pokedexgo/repl"
	"os"
	"time"
)

func main() {
	cacheInterval := 5 * time.Minute
	myCache := pokecache.NewCache(cacheInterval)

	cfg := &repl.Config{
		Commands: map[string]repl.CliCommand{
			"exit": {
				Name:        "exit",
				Description: "Exit the Pokedex",
				Callback:    repl.CommandExit,
			},
			"help": {
				Name:        "help",
				Description: "Displays a help message",
				Callback:    repl.CommandHelp,
			},
			"map": {
				Name:        "map",
				Description: "Displays the names of 20 location areas in the Pokemon world",
				Callback:    internal.CommandMap,
			},
			"mapb": {
				Name:        "mapb",
				Description: "Displays the names of previous 20 location areas in the Pokemon world",
				Callback:    internal.CommandMapB,
			},
		},
		Cache: myCache,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		words := repl.CleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		// Look up the command from cfg.Commands
		if cmd, exists := cfg.Commands[commandName]; exists {
			// Pass the config pointer into the callback
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
