package pokeapi

import (
	"encoding/json"
)

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
