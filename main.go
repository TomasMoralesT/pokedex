package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TomasMoralesT/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(),
	}
	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commands := getCommands()
		command, ok := commands[words[0]]
		if ok {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	words := strings.Fields(lowercased)
	return words
}

func commandExit(cfg *config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, _ ...string) error {
	fmt.Println("Usage:")
	fmt.Println("")

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, _ ...string) error {
	res, err := cfg.pokeapiClient.GetLocationArea(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = &res.Next
	if res.Previous != nil {
		cfg.previousURL = res.Previous
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(cfg *config, _ ...string) error {
	if cfg.previousURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	res, err := cfg.pokeapiClient.GetLocationArea(cfg.previousURL)

	if err != nil {
		return err
	}

	cfg.nextURL = &res.Next
	if res.Previous != nil {
		cfg.previousURL = res.Previous
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location area name")
	}

	locationArea := args[0]

	fmt.Printf("Exploring %s...\n", locationArea)

	location, err := cfg.pokeapiClient.GetLocationAreaByName(locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range location.Pokemon {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the Pokémon in the location area",
			callback:    commandExplore,
		},
	}
}
