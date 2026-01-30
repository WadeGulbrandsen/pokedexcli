package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/WadeGulbrandsen/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapi    pokeapi.Client
	nextLocURL *string
	prevLocURL *string
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		command, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else if err := command.callback(cfg); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
