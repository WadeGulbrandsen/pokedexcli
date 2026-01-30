package pokeapi

import (
	"encoding/json"
)

type LocationArea struct {
	Id                int                `json:"id"`
	Name              *string            `json:"name"`
	Location          NamedAPIResource   `json:"location"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}

func (c *Client) ListLocationAreas(page *string) (NamedAPIResourceList, error) {
	url := baseURL + "/location-area/"
	if page != nil {
		url = *page
	}

	bytes, err := c.Get(url)
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	pokeres := NamedAPIResourceList{}
	if err := json.Unmarshal(bytes, &pokeres); err != nil {
		return NamedAPIResourceList{}, err
	}
	return pokeres, nil
}

func (c *Client) GetLocationArea(area_name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + area_name
	bytes, err := c.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	la := LocationArea{}
	if err := json.Unmarshal(bytes, &la); err != nil {
		return LocationArea{}, err
	}
	return la, nil
}
