package main

import (
	"fmt"
	// "github.com/jinjinir/go-pokedex/internal/pokeapi"
	// "log"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have not caught any pokemon yet")
	} else {
		fmt.Println("Pokemon in Pokedex:")
		for _, pokemon := range cfg.caughtPokemon {
			fmt.Printf("- %v\n", pokemon.Name)
		}
	}
	return nil
}
