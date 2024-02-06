package main

import (
	"fmt"
)

func callbackHelp() {
	fmt.Println("Welcome to the Go-Pokedex! " +
		"Here are all the available options:\n")

	for _, cmd := range availableCommands {
		fmt.Printf("- %v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
