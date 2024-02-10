package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("You have not caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Type/s: %v\n", pokemon.Weight)
	for _, typ := range pokemon.Types {
		fmt.Printf("- %s\n", typ.Type.Name)
	}
	return nil
}
