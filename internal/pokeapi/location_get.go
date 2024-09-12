package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
 areaURL := baseURL + "/location-area/" + locationName

// Adds pokemon to cache 
 if data, found := c.cache.Get(areaURL); found {
	locationResponse := Location{}
	err := json.Unmarshal(data, &locationResponse)
	if err != nil {
		return locationResponse, fmt.Errorf("error unmarshalling data %w", err)
	}
	return locationResponse, nil
 }

 req, err := http.Get(areaURL) 
 if err != nil {
	return Location{}, fmt.Errorf("error getting request: %w", err)
 }
 
 if req.StatusCode > 299 {
	return Location{}, fmt.Errorf("status code error %w", err)
 }

 data, err := io.ReadAll(req.Body)
 if err != nil {
	return Location{}, fmt.Errorf("something went wrong reading result body: %w", err)
 }

 defer req.Body.Close()

 locationResponse := Location{}

 err = json.Unmarshal(data, &locationResponse)
 if err != nil {
	return Location{}, fmt.Errorf("error unmarshalling data: %w", err)
 }

 return locationResponse, nil

}