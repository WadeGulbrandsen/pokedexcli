package pokeapi

import "encoding/json"

type Pokemon struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	BaseXP int           `json:"base_experience"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Stats  []PokemonStat `json:"stats"`
	Types  []PokemonType `json:"types"`
}

type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"type"`
	Slot int              `json:"slot"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	bytes, err := c.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	if err := json.Unmarshal(bytes, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
