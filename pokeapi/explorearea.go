package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (http *Client) ExpoloreArea(locationName string) (pokemons AreaPokemons, err error) {
	url := baseUrl + "/location-area/" + locationName

	data, ok := http.cache.Get(url)
	if ok {
		json.Unmarshal(data, &pokemons)
		fmt.Println("hello", pokemons)
		return
	}

	resp, err := http.client.Get(url)

	if err != nil {
		return
	}

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	json.Unmarshal(byteData, &pokemons)
	return
}
