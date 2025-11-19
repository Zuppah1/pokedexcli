package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonList(name string) (RespPokemonList, error) {

	pageURL := locationURL + name

	if val, ok := c.cache.Get(pageURL); ok {
		pokemonListResp := RespPokemonList{}
		err := json.Unmarshal(val, &pokemonListResp)
		if err != nil {
			return RespPokemonList{}, err
		}

		return pokemonListResp, nil
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return RespPokemonList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonList{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonList{}, nil
	}

	pokemonListResp := RespPokemonList{}

	err = json.Unmarshal(dat, &pokemonListResp)
	if err != nil {
		return RespPokemonList{}, nil
	}

	c.cache.Add(pageURL, dat)

	return pokemonListResp, nil
}
