package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokiUrl string) (pokemon, error) {
	url := baseURL + pokiUrl

	if val, ok := c.cache.Get(url); ok {
		poke := pokemon{}
		if err := json.Unmarshal(val, &poke); err != nil {
			return pokemon{}, err
		}
		fmt.Println("From cache")
		return poke, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon{}, err
	}

	poki := pokemon{}

	if err := json.Unmarshal(data, &poki); err != nil {
		return pokemon{}, err
	}

	c.cache.Add(url, data)
	return poki, nil
}
