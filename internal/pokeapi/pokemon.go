package pokeapi

import "encoding/json"

type Pokemon struct {
	Id     int     `json:"id"`
	Name   *string `json:"name"`
	BaseXP int     `json:"base_experience"`
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
