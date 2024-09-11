package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	name := pokemonName
	pokemonURL := baseURL + "/pokemon/" + name

	if data, found := c.cache.Get(name); found {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(data, &pokemonResponse)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshalling pokemon: %w", err)
		}
		return pokemonResponse, nil
	}

	req, err := http.Get(pokemonURL)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error getting request")
	}

	if req.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("status cod error %d", req.StatusCode)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading request body: %w", err)
	}

	defer req.Body.Close()

	pokemonCatch := Pokemon{}

	err = json.Unmarshal(data, &pokemonCatch)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshalling pokemon: %w", err)
	}

	c.cache.Add(pokemonURL, data)

	return pokemonCatch, nil
}