package pokeapi

import (
	"bytes"
	"github.com/primawk/pokedexgo/internal/pokecache"
	"github.com/primawk/pokedexgo/repl"
	"io"
	"os"
	"testing"
	"time"
)

func setupTestConfig() *repl.Config {
	cacheInterval := 5 * time.Minute
	myCache := pokecache.NewCache(cacheInterval)

	return &repl.Config{
		Pokedex: map[string]repl.Pokemon{},
		Cache:   myCache,
	}
}

func TestCommandCatch_Magikarp(t *testing.T) {
	cfg := setupTestConfig()
	pokemonName := "pikachu" // Magikarp has low base experience (~40), making it easy to catch

	// Try catching it up to 10 times to account for randomness
	caught := false
	var err error
	for i := 0; i < 10; i++ {
		err = CommandCatch(cfg, pokemonName)
		if err != nil {
			t.Fatalf("unexpected error during catch: %v", err)
		}

		if _, exists := cfg.Pokedex[pokemonName]; exists {
			caught = true
			break
		}
	}

	if !caught {
		t.Errorf("failed to catch %s after 10 attempts (statistically unlikely with low base experience)", pokemonName)
	}
}

func TestCommandInspect(t *testing.T) {
	url := "https://pokeapi.co/api/v2/pokemon/pikachu"

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

	// Setup mock config with a caught Pokémon
	cfg := &repl.Config{
		Pokedex: map[string]repl.Pokemon{
			"pikachu": {
				Data: PokemonDetails{
					Name:   "pikachu",
					Height: 4,
					Weight: 60,
					Stats: []StatData{
						{Stat: StatInfo{Name: "hp"}, BaseStat: 35},
						{Stat: StatInfo{Name: "attack"}, BaseStat: 55},
					},
					Types: []TypeData{
						{Type: TypeInfo{Name: "electric"}},
					},
				},
			},
		},
	}

	cases := []struct {
		name           string
		pokemonName    string
		expectedOutput string
	}{
		{
			name:        "Inspect caught pokemon",
			pokemonName: "pikachu",
			expectedOutput: "Name: pikachu\n" +
				"Height: 4\n" +
				"Weight: 60\n" +
				"Stats:\n" +
				" - hp: 35\n" +
				" - attack: 55\n" +
				"Types:\n" +
				" - electric\n",
		},
		{
			name:           "Inspect uncaught pokemon",
			pokemonName:    "mewtwo",
			expectedOutput: "you have not caught that pokemon\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Capture stdout to inspect what fmt.Printf/Println prints
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function (casting cfg to interface/type expected by your function)
			// err := CommandInspect((*repl.Config)(cfg), tc.pokemonName)
			err := CommandInspect(cfg, tc.pokemonName)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Read captured output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			got := buf.String()

			if got != tc.expectedOutput {
				t.Errorf("expected:\n%q\ngot:\n%q", tc.expectedOutput, got)
			}
		})
	}
}
