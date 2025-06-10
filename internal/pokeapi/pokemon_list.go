package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(argument *string) (AreaResponse, error) {
	url := baseURL + "/location-area/"
	if argument != nil {
		url += *argument
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := AreaResponse{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return AreaResponse{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaResponse{}, err
	}

	pokemonResp := AreaResponse{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return AreaResponse{}, err
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}
