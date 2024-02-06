package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Go-Pokedex> ")
		scanner.Scan()
		text := scanner.Text()

		fmt.Println("echoing:", text)
	}
}
