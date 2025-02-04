package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokiUrl string) (Pokemon, error) {
	url := baseURL + pokiUrl

	if val, ok := c.cache.Get(url); ok {
		poke := Pokemon{}
		if err := json.Unmarshal(val, &poke); err != nil {
			return Pokemon{}, err
		}
		fmt.Println("From cache")
		return poke, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	poki := Pokemon{}

	if err := json.Unmarshal(data, &poki); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return poki, nil
}
