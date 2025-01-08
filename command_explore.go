package main

import "fmt"

func commandExplore(cfg *config, l ...string) error {
	if len(l) == 0 {
		return fmt.Errorf("no location provided")
	}

	locationsResp, err := cfg.pokeapiClient.ExploreLocation(l[0])
	if err != nil {
		return err
	}

	for _, pkm := range locationsResp.PokemonEncounters {
		fmt.Println(pkm.Pokemon.Name)
	}

	return nil
}
