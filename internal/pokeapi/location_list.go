package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeLocations, error){
	
	locationUrl := baseURL + "/location-area"
	if pageURL != nil {
		locationUrl = *pageURL
	}

	if data, found := c.cache.Get(locationUrl); found {
		locationList := PokeLocations{}
		err := json.Unmarshal(data, &locationList)
		if err != nil {
			return PokeLocations{}, fmt.Errorf("error unmarshalling data")
		}
		return locationList, nil
	}
	
	res, err := http.Get(locationUrl)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("error with map get request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokeLocations{}, fmt.Errorf("error with status code: %d", res.StatusCode)
	}

	if err != nil {
		return PokeLocations{}, fmt.Errorf("something went wrong reading everything: %w", err)
	}

	data := []byte(body)

	locationList := PokeLocations{}
	err = json.Unmarshal(data, &locationList)
	if err != nil {
		return PokeLocations{}, fmt.Errorf("error mapping: %w", err)
	}
	c.cache.Add(locationUrl, data)
	return locationList, nil

	
}