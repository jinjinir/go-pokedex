package pokeapi

// this package will be responsible for connecting with the pokeapi over the
// internet

import (
	"net/http"
	"time"

	"github.com/jinjinir/go-pokedex/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			// after one minute, the http connection will be killed
			Timeout: time.Minute,
		},
	}
}
