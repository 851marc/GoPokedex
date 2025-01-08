package main

import (
	"MODULE_NAME/internal/pokeapi"
	"time"
)

func main() {
	config := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(config)
}
