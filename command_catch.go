package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, p ...string) error {
	if len(p) == 0 {
		return fmt.Errorf("no pokemon provided")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(p[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	capture := float64(rand.Intn(100)) / float64(pokemon.BaseExperience)

	if capture >= 0.05 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
