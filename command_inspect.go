package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]

	if pokemon, exists := cfg.caughtPokemon[name]; exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		fmt.Printf(" -%s: %d\n", pokemon.Stats[0].Stat.Name, pokemon.Stats[0].BaseStat)
		fmt.Printf(" -%s: %d\n", pokemon.Stats[1].Stat.Name, pokemon.Stats[1].BaseStat)
		fmt.Printf(" -%s: %d\n", pokemon.Stats[2].Stat.Name, pokemon.Stats[2].BaseStat)
		fmt.Printf(" -%s: %d\n", pokemon.Stats[3].Stat.Name, pokemon.Stats[3].BaseStat)
		fmt.Printf(" -%s: %d\n", pokemon.Stats[4].Stat.Name, pokemon.Stats[4].BaseStat)
		fmt.Printf(" -%s: %d\n", pokemon.Stats[5].Stat.Name, pokemon.Stats[5].BaseStat)
		fmt.Println("Types:")
		fmt.Printf(" -%s\n", pokemon.Types[0].Type.Name)
		fmt.Printf(" -%s\n", pokemon.Types[1].Type.Name)

	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
