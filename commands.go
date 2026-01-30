package main

import (
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations on the map",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "List the Pokemon in the given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <name>",
			description: "Try to catch the named Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <name>",
			description: "Inspect the stats of the named Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the Pokemon in your Pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Printf("Your Pokedex: %d caught\n", len(cfg.caughtPokemon))
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("  Your Pokedex is empty. Go catch some Pokemon!")
		return nil
	}
	for name := range cfg.caughtPokemon {
		fmt.Println(" - ", name)
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No name given: Usage 'inspect <name>'")
	}
	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("You need to catch a %s before you can inspect it.", name)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, p_type := range pokemon.Types {
		fmt.Printf(" - %s\n", p_type.Type.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No name given: Usage 'catch <name>'")
	}
	name := args[0]
	pokemon, err := cfg.pokeapi.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if roll := rand.Intn(pokemon.BaseXP); roll < 50 {
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No area name given: Usage 'explore <area_name>'")
	}
	area_name := args[0]
	la, err := cfg.pokeapi.GetLocationArea(area_name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", area_name)
	for _, p := range la.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func getMapResults(cfg *config, endpoint *string) error {
	pokeres, err := cfg.pokeapi.ListLocationAreas(endpoint)
	if err != nil {
		return err
	}
	cfg.nextLocURL = pokeres.Next
	cfg.prevLocURL = pokeres.Previous
	for _, result := range pokeres.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	return getMapResults(cfg, cfg.nextLocURL)
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return getMapResults(cfg, cfg.prevLocURL)
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
