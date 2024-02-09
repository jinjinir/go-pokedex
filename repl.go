package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Go-Pokedex> ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println(
				"Invalid Command. Type 'help' to see available options.")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"clear": {
			name:        "clear",
			description: "Clears the screen",
			callback:    callbackClear,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"explore": {
			name:        "explore {location area}",
			description: "Lists the pokemon availabel in the area",
			callback:    callbackExplore,
		},
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "View information about a caught pokemon",
			callback:    callbackInspect,
		},
		"map": {
			name:        "map",
			description: "Shows the previous map pages",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous map pages",
			callback:    callbackMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows a list of all caught pokemon",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Quits the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
