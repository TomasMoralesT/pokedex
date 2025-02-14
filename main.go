package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/TomasMoralesT/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
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

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("missing Pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := cfg.pokeapiClient.GetPokemonData((args[0]))
	if err != nil {
		return err
	}

	baseRate := 80
	scalingFactor := 4
	threshold := baseRate - (pokemon.BaseExperience / scalingFactor)
	randNum := rand.Intn(100)
	caught := randNum < threshold

	if caught {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("missing Pokemon name")
	}
	value, exists := cfg.caughtPokemon[args[0]]
	if exists {
		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("Height: %d\n", value.Height)
		fmt.Printf("Weight: %d\n", value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range value.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}

	} else {
		fmt.Printf("%s not caught yet!\n", args[0])
	}
	return nil
}

func commandPokedex(cfg *config, _ ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some Pokémon first!")
		return nil // Assuming your function returns an error
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
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
	caughtPokemon map[string]pokeapi.Pokemon
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
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints Pokemon's name, height, weight, stats and type",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
