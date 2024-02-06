package main

import (
	"fmt"
)

func callbackHelp() error {
	fmt.Println("Welcome to the Go-Pokedex! ")
	fmt.Println("Here are all the available options:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf("- %v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
