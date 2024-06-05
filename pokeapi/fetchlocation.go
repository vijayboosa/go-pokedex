package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (http *Client) FetchLocation(url string) (result LocationResult, err error) {

	data, ok := http.cache.Get(url)
	if ok {
		err = json.Unmarshal(data, &result)
		fmt.Println(result)
		return
	}

	res, err := http.client.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)

	if err != nil {
		return
	}

	// add data to cache

	http.cache.Add(url, data)
	err = json.Unmarshal(data, &result)
	return
}
