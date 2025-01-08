package main

import (
	"time"

	"github.com/851marc/GoPokedex/internal/pokeapi"
)

func main() {
	config := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(config)
}
