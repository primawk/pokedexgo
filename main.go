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

	for scanner.Scan() {
		input := repl.CleanInput(scanner.Text())[0]
		cmd, exists := repl.Commands[input]

		if exists {
    		err := cmd.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
    		fmt.Println("Unknown command")
		}


		fmt.Print("Pokedex > ")
	}
}

