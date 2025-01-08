package main

import (
	"fmt"
)

func commandInspect(cfg *config, p ...string) error {
	if len(p) == 0 {
		return fmt.Errorf("no pokemon provided")
	}

	pokemon, ok := cfg.caughtPokemon[p[0]]
	if !ok {
		return fmt.Errorf("pokemon not caught")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  %s\n", t.Type.Name)
	}
	return nil
}
