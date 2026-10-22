package internal

import (
	"encoding/json"
	"fmt"
	"github.com/primawk/pokedexgo/repl"
	"io"
	"math/rand"
	"net/http"
)

type PokemonDetails struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       any    `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  any    `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      any    `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			Ability  any  `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
	} `json:"past_abilities"`
}

func CommandMap(c *repl.Config, _ string) error {
	resource := repl.LocationArea{}

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
	resource := repl.LocationArea{}

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
	resource := repl.LocationAreasResource{}

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
	resource := PokemonDetails{}
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
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
