package repl

import (
    "fmt"
    "os"
	"strings"
)


var Commands map[string]cliCommand

func init() {
    Commands = map[string]cliCommand{
        "exit": {
            Name:        "exit",
            Description: "Exit the Pokedex",
            Callback:    commandExit,
        },
        "help": {
            Name:        "help",
            Description: "Displays a help message",
            Callback:    commandHelp,
        },
    }
}


func CleanInput(text string) []string {
	words := strings.Fields(text)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return words
}


type cliCommand struct {
    Name        string
    Description string
    Callback    func() error
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
    return nil
}

func commandHelp() error {  
	msg := "Welcome to the Pokedex!\nUsage:\n\n"

	for name, cmd := range Commands {
        msg += fmt.Sprintf("%s: %s\n", name, cmd.Description)
    }

	fmt.Print(msg)
	return nil
}

