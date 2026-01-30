package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next string
	prev string
}

type pokeapi_response struct {
	Count    int                  `json:"count"`
	Next     string               `json:"next"`
	Previous string               `json:"previous"`
	Results  []named_api_resource `json:"results"`
}

type named_api_resource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func getPokeAPIResults(cfg *config, endpoint string) ([]named_api_resource, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return []named_api_resource{}, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	pokeres := pokeapi_response{Next: "", Previous: ""}
	if err := decoder.Decode(&pokeres); err != nil {
		return []named_api_resource{}, err
	}
	cfg.next = pokeres.Next
	cfg.prev = pokeres.Previous
	return pokeres.Results, nil
}

func commandMap(cfg *config) error {
	endpoint := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != "" {
		endpoint = cfg.next
	}
	results, err := getPokeAPIResults(cfg, endpoint)
	if err != nil {
		return err
	}
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	results, err := getPokeAPIResults(cfg, cfg.prev)
	if err != nil {
		return err
	}
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
