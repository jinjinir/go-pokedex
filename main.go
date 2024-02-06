package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Go-Pokedex> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("echoing:", text)
}
