package pokeapi

import (
	"net/http"
	"time"

	"github.com/TokiLoshi/pokedexcli/internal/pokecache"
)

// Client

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

// Create new Client
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}