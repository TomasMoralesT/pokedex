package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	words := strings.Fields(lowercased)
	return words
}
