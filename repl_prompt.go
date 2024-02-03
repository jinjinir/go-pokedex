package main

import (
	"fmt"
)

// program name to be printed
var cliName string = "Go Pokedex"

// prompt to be displayed at the start of each loop of the REPL
func printPrompt() {
	fmt.Println("Pokedex> ")
}

// inform the user of invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found. Please refer to the help page for the proper usage.")
}
