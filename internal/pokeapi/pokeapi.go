package pokeapi

// this package will be responsible for connecting with the pokeapi over the
// internet

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			// after one minute, the http connection will be killed
			Timeout: time.Minute,
		},
	}
}
