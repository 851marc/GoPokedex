package main

import (
	"MODULE_NAME/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		cmd, ok := getCommands()[commandName]
		if !ok {
			fmt.Print("Unknown command")
			continue
		}

		err := cmd.callback(config, words[1:]...)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		// fmt.Printf("Your command was: %s\n", commandName)
		//break
	}
}

func cleanInput(text string) []string {
	stringSlice := strings.Fields(strings.TrimSpace(strings.ToLower(text)))

	return stringSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Dispaly 20 maps",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Dispaly previous 20 maps",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display caught pokemon",
			callback:    commandPokedex,
		},
	}
}

// func GetMaps(url string) (PokemonMapResp, error) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := io.ReadAll(res.Body)
// 	res.Body.Close()
// 	if res.StatusCode > 299 {
// 		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp := PokemonMapResp{}
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return resp, err
// }

// type PokemonMapResp struct {
// 	Count    int          `json:"count"`
// 	Next     string       `json:"next"`
// 	Previous string       `json:"previous"`
// 	Results  []PokemonMap `json:"results"`
// }

// type PokemonMap struct {
// 	ID                   int    `json:"id"`
// 	Name                 string `json:"name"`
// 	GameIndex            int    `json:"game_index"`
// 	EncounterMethodRates []struct {
// 		EncounterMethod struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"encounter_method"`
// 		VersionDetails []struct {
// 			Rate    int `json:"rate"`
// 			Version struct {
// 				Name string `json:"name"`
// 				URL  string `json:"url"`
// 			} `json:"version"`
// 		} `json:"version_details"`
// 	} `json:"encounter_method_rates"`
// 	Location struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"location"`
// 	Names []struct {
// 		Name     string `json:"name"`
// 		Language struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"language"`
// 	} `json:"names"`
// 	PokemonEncounters []struct {
// 		Pokemon struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"pokemon"`
// 		VersionDetails []struct {
// 			Version struct {
// 				Name string `json:"name"`
// 				URL  string `json:"url"`
// 			} `json:"version"`
// 			MaxChance        int `json:"max_chance"`
// 			EncounterDetails []struct {
// 				MinLevel        int   `json:"min_level"`
// 				MaxLevel        int   `json:"max_level"`
// 				ConditionValues []any `json:"condition_values"`
// 				Chance          int   `json:"chance"`
// 				Method          struct {
// 					Name string `json:"name"`
// 					URL  string `json:"url"`
// 				} `json:"method"`
// 			} `json:"encounter_details"`
// 		} `json:"version_details"`
// 	} `json:"pokemon_encounters"`
// }
