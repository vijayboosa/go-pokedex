package pokeapi

import (
	"encoding/json"
	"io"
)

func (http *Client) CatchPokemon(pokemonName string) (pokemon Pokemon, err error) {

	url := baseUrl + "/pokemon/" + pokemonName

	if data, ok := http.cache.Get(url); ok {
		json.Unmarshal(data, &pokemon)
		return
	}

	resp, err := http.client.Get(url)
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	http.cache.Add(url, body)

	err = json.Unmarshal(body, &pokemon)

	return
}
