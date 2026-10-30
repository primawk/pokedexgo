package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/primawk/pokedexgo/repl"
	"io"
	"math/rand"
	"net/http"
)

func CommandMap(c *repl.Config, _ string) error {
	resource := LocationArea{}

	url := c.Next

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	if value, ok := c.Cache.Get(url); ok {
		if err := json.Unmarshal(value, &resource); err != nil {
			return fmt.Errorf("Error unmarshalling JSON: %v", err)
		}

		for _, loc := range resource.Results {
			fmt.Printf("%s\n", loc.Name)
		}

		c.Next = resource.Next
		c.Previous = resource.Previous

		return nil
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return err
	}

	c.Cache.Add(url, body)

	if err := json.Unmarshal(body, &resource); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	c.Next = resource.Next
	c.Previous = resource.Previous

	for _, loc := range resource.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

func CommandMapB(c *repl.Config, _ string) error {
	resource := LocationArea{}

	url := c.Previous

	if url == "" {
		return fmt.Errorf("You are on the first page")
	}

	if value, ok := c.Cache.Get(url); ok {
		if err := json.Unmarshal(value, &resource); err != nil {
			return fmt.Errorf("Error unmarshalling JSON: %v", err)
		}

		for _, loc := range resource.Results {
			fmt.Printf("%s\n", loc.Name)
		}

		c.Next = resource.Next
		c.Previous = resource.Previous

		return nil
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &resource); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	c.Next = resource.Next
	c.Previous = resource.Previous

	for _, loc := range resource.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

func CommandExplore(c *repl.Config, name string) error {
	resource := LocationAreasResource{}

	url := "https://pokeapi.co/api/v2/location-area/" + name

	fmt.Printf("exploring %s...\n", name)

	if value, ok := c.Cache.Get(url); ok {
		fmt.Println("retrieve from cached data")
		if err := json.Unmarshal(value, &resource); err != nil {
			return fmt.Errorf("Error unmarshalling JSON: %v", err)
		}

		for _, pokemon := range resource.PokemonEncounters {
			fmt.Printf("%s\n", pokemon.Pokemon.Name)
		}

		return nil
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return err
	}

	c.Cache.Add(url, body)

	if err := json.Unmarshal(body, &resource); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range resource.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func CommandCatch(c *repl.Config, name string) error {
	if _, exists := c.Pokedex[name]; exists {
		fmt.Println("You already caught that pokemon")
		return nil
	}
	resource := repl.PokemonDetails{}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &resource); err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	point := rand.Intn(resource.BaseExperience + 1)

	if point > resource.BaseExperience/2 {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		c.Pokedex[name] = repl.Pokemon{
			Data: resource,
		}
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func CommandInspect(c *repl.Config, name string) error {
	if value, exists := c.Pokedex[name]; exists {
		fmt.Printf("Name: %s\n", value.Data.Name)
		fmt.Printf("Height: %d\n", value.Data.Height)
		fmt.Printf("Weight: %d\n", value.Data.Weight)
		fmt.Println("Stats:")

		for _, stat := range value.Data.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")

		for _, item := range value.Data.Types {
			fmt.Printf(" - %s\n", item.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}

func CommandPokedex(c *repl.Config, _ string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("You don't have any pokemon in your pokedex.")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, item := range c.Pokedex {
		fmt.Printf(" - %s\n", item.Data.Name)
	}
	return nil
}
