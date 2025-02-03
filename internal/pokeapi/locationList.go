package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (locationResponse, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locRes := locationResponse{}
		err := json.Unmarshal(val, &locRes)
		if err != nil {
			return locationResponse{}, err
		}
		return locRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationResponse{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return locationResponse{}, err
	}

	locRes := locationResponse{}

	if err := json.Unmarshal(dat, &locRes); err != nil {
		return locationResponse{}, err
	}
	c.cache.Add(url, dat)
	return locRes, nil
}
