package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/WadeGulbrandsen/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	key := "GET: " + url
	if bytes, ok := c.cache.Get(key); ok {
		return bytes, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error %d getting %s: %s", res.StatusCode, url, res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(key, bytes)
	return bytes, nil
}
