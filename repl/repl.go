package repl

import (
	"fmt"
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
	// You can add other state fields here later (e.g., NextURL string)
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
