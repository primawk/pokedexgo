package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/primawk/pokedexgo/repl"
)

func CommandMap(c *repl.Config) error {
	resource := repl.LocationArea{}
	url := c.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
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

func CommandMapB(c *repl.Config) error {
	resource := repl.LocationArea{}
	url := c.Previous
	if url == "" {
		return fmt.Errorf("You are on the first page")
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
