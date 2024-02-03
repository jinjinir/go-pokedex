package main

import (
	"fmt"
	"os"
	"os/exec"
)

// displayHelp shows the user some hardcoded functions
func displayHelp() {
	fmt.Printf(
		"%v! Here is a quick overview of the available commands!\n", cliName,
	)
	fmt.Println("help\t-\tShow available commands.")
	fmt.Println("clear\t-\tClears the screen.")
	fmt.Println("exit\t-\tExits the program.")
}

// clearScreen clears the REPL screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
