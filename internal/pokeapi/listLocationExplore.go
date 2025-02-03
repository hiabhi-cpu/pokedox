package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationExplore(pageUrl string) (locExploreResponse, error) {
	url := baseURL + pageUrl
	// fmt.Println(url)

	if val, ok := c.cache.Get(url); ok {
		locExploreRes := locExploreResponse{}
		err := json.Unmarshal(val, &locExploreRes)
		if err != nil {
			return locExploreResponse{}, err
		}
		fmt.Println("from cache")
		return locExploreRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locExploreResponse{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locExploreResponse{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locExploreResponse{}, nil
	}

	locExpRes := locExploreResponse{}

	if err := json.Unmarshal(data, &locExpRes); err != nil {
		return locExploreResponse{}, nil
	}
	c.cache.Add(url, data)
	return locExpRes, nil
}
