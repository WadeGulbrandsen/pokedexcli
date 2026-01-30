package main

import (
	"time"

	"github.com/WadeGulbrandsen/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapi:       pokeapi.NewClient(5*time.Second, 7*time.Second),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startREPL(&cfg)
}
