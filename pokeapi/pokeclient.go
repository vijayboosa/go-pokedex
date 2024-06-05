package pokeapi

import (
	"net/http"
	"time"

	"github.com/vijayboosa/pokedex-cli/pokecache"
)

type LocationResult struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Area string `json:"name"`
	} `json:"results"`
}

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func NewClient() Client {
	cache := pokecache.NewCache(time.Second * 20)
	client := &http.Client{
		Timeout: time.Second * 6,
	}

	return Client{
		client: *client,
		cache:  cache,
	}
}
