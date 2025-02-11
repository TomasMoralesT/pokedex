package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		lowercased := strings.ToLower(input)
		words := strings.Fields(lowercased)
		if len(words) > 0 {
			fmt.Printf("Your command was: %s\n", words[0])

		}
	}
}

func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	words := strings.Fields(lowercased)
	return words
}
