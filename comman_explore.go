package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location area")
	}

	name := args[0]

	pokemonListResp, err := cfg.pokeapiClient.PokemonList(name)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s", name)
	fmt.Println()
	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemonListResp.PokemonEncounters {
		fmt.Printf("- %s", pokemon.Pokemon.Name)
		fmt.Println()
	}
	fmt.Println()

	return nil
}
