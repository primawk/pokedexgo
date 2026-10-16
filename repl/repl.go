package repl

import (
	"fmt"
	"github.com/primawk/pokedexgo/internal/pokecache"
	"os"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

type Config struct {
	Commands map[string]CliCommand
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}

func CommandExit(c *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func CommandHelp(c *Config) error {
	msg := "Welcome to the Pokedex!\nUsage:\n\n"

	// Access the commands map stored inside the Config instance
	for name, cmd := range c.Commands {
		msg += fmt.Sprintf("%s: %s\n", name, cmd.Description)
	}

	fmt.Print(msg)
	return nil
}
